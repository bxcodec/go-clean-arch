package article

import (
	"context"

	"github.com/bxcodec/go-clean-arch/v2/models"
)

type ArticleRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]*models.Article, error)
	GetByID(ctx context.Context, id int64) (*models.Article, error)
	GetByTitle(ctx context.Context, title string) (*models.Article, error)
	Update(ctx context.Context, article *models.Article) (*models.Article, error)
	Store(ctx context.Context, a *models.Article) (int64, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
