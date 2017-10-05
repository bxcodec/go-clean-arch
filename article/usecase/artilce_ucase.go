package article

import (
	"strconv"
	"time"

	"github.com/bxcodec/go-clean-arch/article"
	"github.com/bxcodec/go-clean-arch/article/repository"
)

type ArticleUsecase interface {
	Fetch(cursor string, num int64) ([]*article.Article, string, error)
	GetByID(id int64) (*article.Article, error)
	Update(ar *article.Article) (*article.Article, error)
	GetByTitle(title string) (*article.Article, error)
	Store(*article.Article) (*article.Article, error)
	Delete(id int64) (bool, error)
}

type articleUsecase struct {
	articleRepos repository.ArticleRepository
}

func (a *articleUsecase) Fetch(cursor string, num int64) ([]*article.Article, string, error) {
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

func (a *articleUsecase) GetByID(id int64) (*article.Article, error) {

	return a.articleRepos.GetByID(id)
}

func (a *articleUsecase) Update(ar *article.Article) (*article.Article, error) {
	ar.UpdatedAt = time.Now()
	return a.articleRepos.Update(ar)
}

func (a *articleUsecase) GetByTitle(title string) (*article.Article, error) {

	return a.articleRepos.GetByTitle(title)
}

func (a *articleUsecase) Store(m *article.Article) (*article.Article, error) {

	existedArticle, _ := a.GetByTitle(m.Title)
	if existedArticle != nil {
		return nil, article.CONFLIT_ERROR
	}

	id, err := a.articleRepos.Store(m)
	if err != nil {
		return nil, err
	}

	m.ID = id
	return m, nil
}

func (a *articleUsecase) Delete(id int64) (bool, error) {
	existedArticle, _ := a.GetByID(id)

	if existedArticle == nil {
		return false, article.NOT_FOUND_ERROR
	}

	return a.articleRepos.Delete(id)
}

func NewArticleUsecase(a repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{a}
}
