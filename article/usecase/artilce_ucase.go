package usecase

import (
	"strconv"
	"time"

	"github.com/bxcodec/go-clean-arch/models"
	"github.com/sirupsen/logrus"

	"github.com/bxcodec/go-clean-arch/article"
	_authorRepo "github.com/bxcodec/go-clean-arch/author"
)

type articleUsecase struct {
	articleRepos article.ArticleRepository
	authorRepo   _authorRepo.AuthorRepository
}

type authorChanel struct {
	Author *models.Author
	Error  error
}

func NewArticleUsecase(a article.ArticleRepository, ar _authorRepo.AuthorRepository) article.ArticleUsecase {
	return &articleUsecase{
		articleRepos: a,
		authorRepo:   ar,
	}
}

func (a *articleUsecase) getAuthorDetail(item *models.Article, authorChan chan authorChanel) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Debug("Recovered in ", r)
		}
	}()

	res, err := a.authorRepo.GetByID(item.Author.ID)
	holder := authorChanel{
		Author: res,
		Error:  err,
	}
	authorChan <- holder
}
func (a *articleUsecase) getAuthorDetails(data []*models.Article) ([]*models.Article, error) {
	chAuthor := make(chan authorChanel)
	defer close(chAuthor)
	existingAuthorMap := make(map[int64]bool)
	totalCall := 0
	for _, item := range data {
		if _, ok := existingAuthorMap[item.Author.ID]; !ok {
			existingAuthorMap[item.Author.ID] = true
			go a.getAuthorDetail(item, chAuthor)
		}
		totalCall++
	}

	mapAuthor := make(map[int64]*models.Author)
	totalGorutineCalled := len(existingAuthorMap)
	for i := 0; i < totalGorutineCalled; i++ {
		select {
		case a := <-chAuthor:
			if a.Error == nil && a.Author != nil {
				mapAuthor[a.Author.ID] = a.Author
			}
		case <-time.After(time.Second * 1):
			logrus.Warn("Timeout when calling user detail")

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

func (a *articleUsecase) Fetch(cursor string, num int64) ([]*models.Article, string, error) {
	if num == 0 {
		num = 10
	}

	listArticle, err := a.articleRepos.Fetch(cursor, num)
	if err != nil {
		return nil, "", err
	}

	nextCursor := ""

	listArticle, err = a.getAuthorDetails(listArticle)
	if err != nil {
		return nil, "", err
	}

	if size := len(listArticle); size == int(num) {
		lastId := listArticle[num-1].ID
		nextCursor = strconv.Itoa(int(lastId))
	}

	return listArticle, nextCursor, nil
}

func (a *articleUsecase) GetByID(id int64) (*models.Article, error) {

	res, err := a.articleRepos.GetByID(id)
	if err != nil {
		return nil, err
	}

	resAuthor, err := a.authorRepo.GetByID(res.Author.ID)
	if err != nil {
		return nil, err
	}
	res.Author = *resAuthor
	return res, nil
}

func (a *articleUsecase) Update(ar *models.Article) (*models.Article, error) {
	ar.UpdatedAt = time.Now()
	return a.articleRepos.Update(ar)
}

func (a *articleUsecase) GetByTitle(title string) (*models.Article, error) {

	res, err := a.articleRepos.GetByTitle(title)
	if err != nil {
		return nil, err
	}

	resAuthor, err := a.authorRepo.GetByID(res.Author.ID)
	if err != nil {
		return nil, err
	}
	res.Author = *resAuthor

	return res, nil
}

func (a *articleUsecase) Store(m *models.Article) (*models.Article, error) {

	existedArticle, _ := a.GetByTitle(m.Title)
	if existedArticle != nil {
		return nil, models.CONFLIT_ERROR
	}

	id, err := a.articleRepos.Store(m)
	if err != nil {
		return nil, err
	}

	m.ID = id
	return m, nil
}

func (a *articleUsecase) Delete(id int64) (bool, error) {
	existedArticle, _ := a.articleRepos.GetByID(id)
	logrus.Info("Masuk Sini")
	if existedArticle == nil {
		logrus.Info("Masuk Sini2")
		return false, models.NOT_FOUND_ERROR
	}
	logrus.Info("Masuk Sini3")

	return a.articleRepos.Delete(id)
}
