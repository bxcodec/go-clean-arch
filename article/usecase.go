package article

import (
	"context"

	model "github.com/bxcodec/go-clean-arch/v2/models"
)

type ArticleUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]*model.Article, string, error)
	GetByID(ctx context.Context, id int64) (*model.Article, error)
	Update(ctx context.Context, ar *model.Article) (*model.Article, error)
	GetByTitle(ctx context.Context, title string) (*model.Article, error)
	Store(context.Context, *model.Article) (*model.Article, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
