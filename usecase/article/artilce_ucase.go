package article

import (
	"strconv"

	"github.com/bxcodec/go-clean-arch/models"
	"github.com/bxcodec/go-clean-arch/repository"
	"github.com/bxcodec/go-clean-arch/usecase"
)

type articleUsecase struct {
	articleRepos repository.ArticleRepository
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

	if size := len(listArticle); size == int(num) {
		lastId := listArticle[num-1].ID
		nextCursor = strconv.Itoa(int(lastId))
	}

	return listArticle, nextCursor, nil
}

func (a *articleUsecase) GetByID(id int64) (*models.Article, error) {

	return a.articleRepos.GetByID(id)
}
func (a *articleUsecase) GetByTitle(title string) (*models.Article, error) {

	return a.articleRepos.GetByTitle(title)
}

func (a *articleUsecase) Store(m *models.Article) (*models.Article, error) {

	existedArticle, _ := a.GetByTitle(m.Title)
	if existedArticle != nil {
		return nil, models.NewErrorConflict()
	}

	id, err := a.articleRepos.Store(m)
	if err != nil {
		return nil, err
	}

	m.ID = id
	return m, nil
}

func NewArticleUsecase(a repository.ArticleRepository) usecase.ArticleUsecase {
	return &articleUsecase{a}
}
