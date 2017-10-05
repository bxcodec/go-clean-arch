package repository

import models "github.com/bxcodec/go-clean-arch/article"

type ArticleRepository interface {
	Fetch(cursor string, num int64) ([]*models.Article, error)
	GetByID(id int64) (*models.Article, error)
	GetByTitle(title string) (*models.Article, error)
	Update(article *models.Article) (*models.Article, error)
	Store(a *models.Article) (int64, error)
	Delete(id int64) (bool, error)
}
