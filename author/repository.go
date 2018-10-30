package author

import (
	"context"

	"github.com/bxcodec/go-clean-arch/v2/models"
)

type AuthorRepository interface {
	GetByID(ctx context.Context, id int64) (*models.Author, error)
}
