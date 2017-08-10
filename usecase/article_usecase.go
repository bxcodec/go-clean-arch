package usecase

import "github.com/bxcodec/go-clean-arch/models"

type ArticleUsecase interface {
	Fetch(cursor string, num int64) ([]*models.Article, string, error)
	GetByID(id int64) (*models.Article, error)
	Update(ar *models.Article) (*models.Article, error)
	GetByTitle(title string) (*models.Article, error)
	Store(*models.Article) (*models.Article, error)
	Delete(id int64) (bool, error)
}
