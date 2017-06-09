package repository

import "github.com/bxcodec/go-clean-arch/models"

type ArticleRepository interface {
	Fetch(cursor string, num int64) ([]*models.Article, error)
}
type CategoryRepository interface {
	Fetch(articleID int64) ([]*models.Category, error)
}
