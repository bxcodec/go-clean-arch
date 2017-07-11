package category

import (
	"strconv"

	"github.com/bxcodec/go-clean-arch/models"
	"github.com/bxcodec/go-clean-arch/repository"
	"github.com/bxcodec/go-clean-arch/usecase"
)

type categoryUsecase struct {
	categoryRepos repository.CategoryRepository
}

func (a *categoryUsecase) Fetch(cursor string, num int64) ([]*models.Category, string, error) {
	if num == 0 {
		num = 10
	}

	listCategory, err := a.categoryRepos.Fetch(cursor, num)
	if err != nil {
		return nil, "", err
	}
	nextCursor := ""

	if size := len(listCategory); size == int(num) {
		lastId := listCategory[num-1].ID
		nextCursor = strconv.Itoa(int(lastId))
	}

	return listCategory, nextCursor, nil
}

func (a *categoryUsecase) GetByID(id int64) (*models.Category, error) {

	return a.categoryRepos.GetByID(id)
}
func (a *categoryUsecase) GetByName(title string) (*models.Category, error) {

	return a.categoryRepos.GetByName(title)
}

func (a *categoryUsecase) Store(m *models.Category) (*models.Category, error) {

	existedCategory, _ := a.GetByName(m.Name)
	if existedCategory != nil {
		return nil, models.NewErrorConflict()
	}

	id, err := a.categoryRepos.Store(m)
	if err != nil {
		return nil, err
	}

	m.ID = id
	return m, nil
}

func (a *categoryUsecase) Delete(id int64) (bool, error) {
	existedCategory, _ := a.GetByID(id)

	if existedCategory == nil {
		return false, models.NewErrorNotFound()
	}

	return a.categoryRepos.Delete(id)
}

func NewCategoryUsecase(a repository.CategoryRepository) usecase.CategoryUsecase {
	return &categoryUsecase{a}
}
