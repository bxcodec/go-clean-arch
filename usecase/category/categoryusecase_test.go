package category_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/bxcodec/go-clean-arch/models"
	"github.com/bxcodec/go-clean-arch/repository/mocks"
	ucase "github.com/bxcodec/go-clean-arch/usecase/category"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	mockCategoryRepo := new(mocks.CategoryRepository)
	var mockCategory models.Category
	err := faker.FakeData(&mockCategory)
	assert.NoError(t, err)

	mockListArtilce := make([]*models.Category, 0)
	mockListArtilce = append(mockListArtilce, &mockCategory)
	mockCategoryRepo.On("Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(mockListArtilce, nil)
	u := ucase.NewCategoryUsecase(mockCategoryRepo)
	num := int64(1)
	cursor := "12"
	list, nextCursor, err := u.Fetch(cursor, num)
	cursorExpected := strconv.Itoa(int(mockCategory.ID))
	assert.Equal(t, cursorExpected, nextCursor)
	assert.NotEmpty(t, nextCursor)
	assert.NoError(t, err)
	assert.Len(t, list, len(mockListArtilce))

	mockCategoryRepo.AssertCalled(t, "Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64"))

}

func TestFetchError(t *testing.T) {
	mockCategoryRepo := new(mocks.CategoryRepository)

	mockCategoryRepo.On("Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpexted Error"))
	u := ucase.NewCategoryUsecase(mockCategoryRepo)
	num := int64(1)
	cursor := "12"
	list, nextCursor, err := u.Fetch(cursor, num)

	assert.Empty(t, nextCursor)
	assert.Error(t, err)
	assert.Len(t, list, 0)
	mockCategoryRepo.AssertCalled(t, "Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64"))

}

func TestGetByID(t *testing.T) {
	mockCategoryRepo := new(mocks.CategoryRepository)
	var mockCategory models.Category
	err := faker.FakeData(&mockCategory)
	assert.NoError(t, err)

	mockCategoryRepo.On("GetByID", mock.AnythingOfType("int64")).Return(&mockCategory, nil)
	defer mockCategoryRepo.AssertCalled(t, "GetByID", mock.AnythingOfType("int64"))
	u := ucase.NewCategoryUsecase(mockCategoryRepo)

	a, err := u.GetByID(mockCategory.ID)

	assert.NoError(t, err)
	assert.NotNil(t, a)

}

func TestStore(t *testing.T) {
	mockCategoryRepo := new(mocks.CategoryRepository)
	var mockCategory models.Category
	err := faker.FakeData(&mockCategory)
	assert.NoError(t, err)
	//set to 0 because this is test from Client, and ID is an AutoIncreament
	tempMockCategory := mockCategory
	tempMockCategory.ID = 0

	mockCategoryRepo.On("GetByName", mock.AnythingOfType("string")).Return(nil, models.NewErrorNotFound())
	mockCategoryRepo.On("Store", mock.AnythingOfType("*models.Category")).Return(mockCategory.ID, nil)
	defer mockCategoryRepo.AssertCalled(t, "GetByName", mock.AnythingOfType("string"))
	defer mockCategoryRepo.AssertCalled(t, "Store", mock.AnythingOfType("*models.Category"))

	u := ucase.NewCategoryUsecase(mockCategoryRepo)

	a, err := u.Store(&tempMockCategory)

	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, mockCategory.Name, tempMockCategory.Name)

}

func TestDelete(t *testing.T) {
	mockCategoryRepo := new(mocks.CategoryRepository)
	var mockCategory models.Category
	err := faker.FakeData(&mockCategory)
	assert.NoError(t, err)

	mockCategoryRepo.On("GetByID", mock.AnythingOfType("int64")).Return(&mockCategory, models.NewErrorNotFound())
	defer mockCategoryRepo.AssertCalled(t, "GetByID", mock.AnythingOfType("int64"))

	mockCategoryRepo.On("Delete", mock.AnythingOfType("int64")).Return(true, nil)
	defer mockCategoryRepo.AssertCalled(t, "Delete", mock.AnythingOfType("int64"))

	u := ucase.NewCategoryUsecase(mockCategoryRepo)

	a, err := u.Delete(mockCategory.ID)

	assert.NoError(t, err)
	assert.True(t, a)

}
