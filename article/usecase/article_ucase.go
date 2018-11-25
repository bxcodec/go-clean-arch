package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/bxcodec/go-clean-arch/models"

	"github.com/bxcodec/go-clean-arch/article"
	_authorRepo "github.com/bxcodec/go-clean-arch/author"
	"golang.org/x/sync/errgroup"
)

type articleUsecase struct {
	articleRepo    article.ArticleRepository
	authorRepo     _authorRepo.AuthorRepository
	contextTimeout time.Duration
}

func NewArticleUsecase(a article.ArticleRepository, ar _authorRepo.AuthorRepository, timeout time.Duration) article.ArticleUsecase {
	return &articleUsecase{
		articleRepo:    a,
		authorRepo:     ar,
		contextTimeout: timeout,
	}
}

/*
* In this function below, I'm using errgroup with the pipeline pattern
* Look how this works in this package explanation
* in godoc: https://godoc.org/golang.org/x/sync/errgroup#ex-Group--Pipeline
 */
func (a *articleUsecase) getAuthorDetails(c context.Context, data []*models.Article) ([]*models.Article, error) {

	g, ctx := errgroup.WithContext(c)

	// Get the author's id
	mapAuthors := map[int64]models.Author{}

	for _, article := range data {
		mapAuthors[article.Author.ID] = models.Author{}
	}
	chanAuthor := make(chan *models.Author)
	for authorID, _ := range mapAuthors {
		authorID := authorID
		g.Go(func() error {
			res, err := a.authorRepo.GetByID(ctx, authorID)
			if err != nil {
				return err
			}
			chanAuthor <- res
			return nil
		})
	}

	go func() {
		g.Wait()
		close(chanAuthor)
	}()

	for author := range chanAuthor {
		if author != nil {
			mapAuthors[author.ID] = *author
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	// merge the author
	for index, item := range data {
		if a, ok := mapAuthors[item.Author.ID]; ok {
			data[index].Author = a
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

	listArticle, err := a.articleRepo.Fetch(ctx, cursor, num)
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

	res, err := a.articleRepo.GetByID(ctx, id)
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
	return a.articleRepo.Update(ctx, ar)
}

func (a *articleUsecase) GetByTitle(c context.Context, title string) (*models.Article, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	res, err := a.articleRepo.GetByTitle(ctx, title)
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
		return nil, models.CONFLICT_ERROR
	}

	id, err := a.articleRepo.Store(ctx, m)
	if err != nil {
		return nil, err
	}

	m.ID = id
	return m, nil
}

func (a *articleUsecase) Delete(c context.Context, id int64) (bool, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedArticle, _ := a.articleRepo.GetByID(ctx, id)
	if existedArticle == nil {
		return false, models.NOT_FOUND_ERROR
	}
	return a.articleRepo.Delete(ctx, id)
}
