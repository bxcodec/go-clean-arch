package repository

import "github.com/bxcodec/go-clean-arch/models"

type ArticleRepository interface {
	Fetch(cursor string, num int64) ([]*models.Article, error)
	GetByID(id int64) (*models.Article, error)
	GetByTitle(title string) (*models.Article, error)
	Store(a *models.Article) (int64, error)
}
type CategoryRepository interface {
	Fetch(articleID int64) ([]*models.Category, error)
}
