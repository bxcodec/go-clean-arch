package article

import (
	model "github.com/bxcodec/go-clean-arch/models"
)

type ArticleUsecase interface {
	Fetch(cursor string, num int64) ([]*model.Article, string, error)
	GetByID(id int64) (*model.Article, error)
	Update(ar *model.Article) (*model.Article, error)
	GetByTitle(title string) (*model.Article, error)
	Store(*model.Article) (*model.Article, error)
	Delete(id int64) (bool, error)
}
