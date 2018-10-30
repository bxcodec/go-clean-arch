package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/bxcodec/go-clean-arch/v2/models"
	"github.com/sirupsen/logrus"

	"github.com/bxcodec/go-clean-arch/v2/article"
	_authorRepo "github.com/bxcodec/go-clean-arch/v2/author"
)

type articleUsecase struct {
	articleRepos   article.ArticleRepository
	authorRepo     _authorRepo.AuthorRepository
	contextTimeout time.Duration
}

type authorChanel struct {
	Author *models.Author
	Error  error
}

func NewArticleUsecase(a article.ArticleRepository, ar _authorRepo.AuthorRepository, timeout time.Duration) article.ArticleUsecase {
	return &articleUsecase{
		articleRepos:   a,
		authorRepo:     ar,
		contextTimeout: timeout,
	}
}

func (a *articleUsecase) getAuthorDetail(ctx context.Context, item *models.Article, authorChan chan authorChanel) {

	res, err := a.authorRepo.GetByID(ctx, item.Author.ID)
	holder := authorChanel{
		Author: res,
		Error:  err,
	}
	if ctx.Err() != nil {
		return // To avoid send on closed channel
	}
	authorChan <- holder
}
func (a *articleUsecase) getAuthorDetails(ctx context.Context, data []*models.Article) ([]*models.Article, error) {
	chAuthor := make(chan authorChanel)
	defer close(chAuthor)
	existingAuthorMap := make(map[int64]bool)
	for _, item := range data {
		if _, ok := existingAuthorMap[item.Author.ID]; !ok {
			existingAuthorMap[item.Author.ID] = true
			go a.getAuthorDetail(ctx, item, chAuthor)
		}

	}

	mapAuthor := make(map[int64]*models.Author)
	totalGorutineCalled := len(existingAuthorMap)
	for i := 0; i < totalGorutineCalled; i++ {
		select {
		case a := <-chAuthor:
			if a.Error == nil {
				if a.Author != nil {
					mapAuthor[a.Author.ID] = a.Author
				}
			} else {
				return nil, a.Error
			}

		case <-ctx.Done():
			logrus.Warn("Timeout when calling user detail")
			return nil, ctx.Err()
		}
	}

	// merge the author
	for index, item := range data {
		if a, ok := mapAuthor[item.Author.ID]; ok {
			data[index].Author = *a
		}
	}
	return data, nil
}

func (a *articleUsecase) Fetch(c context.Context, cursor string, num int64) ([]*models.Article, string, error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listArticle, err := a.articleRepos.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}

	nextCursor := ""

	listArticle, err = a.getAuthorDetails(ctx, listArticle)
	if err != nil {
		return nil, "", err
	}

	if size := len(listArticle); size == int(num) {
		lastId := listArticle[num-1].ID
		nextCursor = strconv.Itoa(int(lastId))
	}

	return listArticle, nextCursor, nil
}

func (a *articleUsecase) GetByID(c context.Context, id int64) (*models.Article, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err := a.articleRepos.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resAuthor, err := a.authorRepo.GetByID(ctx, res.Author.ID)
	if err != nil {
		return nil, err
	}
	res.Author = *resAuthor
	return res, nil
}

func (a *articleUsecase) Update(c context.Context, ar *models.Article) (*models.Article, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	ar.UpdatedAt = time.Now()
	return a.articleRepos.Update(ctx, ar)
}

func (a *articleUsecase) GetByTitle(c context.Context, title string) (*models.Article, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	res, err := a.articleRepos.GetByTitle(ctx, title)
	if err != nil {
		return nil, err
	}

	resAuthor, err := a.authorRepo.GetByID(ctx, res.Author.ID)
	if err != nil {
		return nil, err
	}
	res.Author = *resAuthor

	return res, nil
}

func (a *articleUsecase) Store(c context.Context, m *models.Article) (*models.Article, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedArticle, _ := a.GetByTitle(ctx, m.Title)
	if existedArticle != nil {
		return nil, models.CONFLIT_ERROR
	}

	id, err := a.articleRepos.Store(ctx, m)
	if err != nil {
		return nil, err
	}

	m.ID = id
	return m, nil
}

func (a *articleUsecase) Delete(c context.Context, id int64) (bool, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedArticle, _ := a.articleRepos.GetByID(ctx, id)
	if existedArticle == nil {
		return false, models.NOT_FOUND_ERROR
	}
	return a.articleRepos.Delete(ctx, id)
}
