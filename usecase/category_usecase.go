package usecase

import "github.com/bxcodec/go-clean-arch/models"

type CategoryUsecase interface {
	Fetch(cursor string, num int64) ([]*models.Category, string, error)
	GetByID(id int64) (*models.Category, error)
	GetByName(title string) (*models.Category, error)
	Store(*models.Category) (*models.Category, error)
	Delete(id int64) (bool, error)
}
