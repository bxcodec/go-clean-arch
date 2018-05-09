package author

import (
	"context"

	"github.com/bxcodec/go-clean-arch/models"
)

type AuthorRepository interface {
	GetByID(ctx context.Context, id int64) (*models.Author, error)
}
