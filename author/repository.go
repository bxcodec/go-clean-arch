package author

import "github.com/bxcodec/go-clean-arch/models"

type AuthorRepository interface {
	GetByID(id int64) (*models.Author, error)
}
