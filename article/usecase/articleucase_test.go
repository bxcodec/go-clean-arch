package article_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/bxcodec/faker"
	models "github.com/bxcodec/go-clean-arch/article"
	"github.com/bxcodec/go-clean-arch/article/repository/mocks"
	ucase "github.com/bxcodec/go-clean-arch/article/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	mockArticleRepo := new(mocks.ArticleRepository)
	var mockArticle models.Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)

	mockListArtilce := make([]*models.Article, 0)
	mockListArtilce = append(mockListArtilce, &mockArticle)
	mockArticleRepo.On("Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(mockListArtilce, nil)
	u := ucase.NewArticleUsecase(mockArticleRepo)
	num := int64(1)
	cursor := "12"
	list, nextCursor, err := u.Fetch(cursor, num)
	cursorExpected := strconv.Itoa(int(mockArticle.ID))
	assert.Equal(t, cursorExpected, nextCursor)
	assert.NotEmpty(t, nextCursor)
	assert.NoError(t, err)
	assert.Len(t, list, len(mockListArtilce))

	mockArticleRepo.AssertCalled(t, "Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64"))

}

func TestFetchError(t *testing.T) {
	mockArticleRepo := new(mocks.ArticleRepository)

	mockArticleRepo.On("Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpexted Error"))
	u := ucase.NewArticleUsecase(mockArticleRepo)
	num := int64(1)
	cursor := "12"
	list, nextCursor, err := u.Fetch(cursor, num)

	assert.Empty(t, nextCursor)
	assert.Error(t, err)
	assert.Len(t, list, 0)
	mockArticleRepo.AssertCalled(t, "Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64"))

}

func TestGetByID(t *testing.T) {
	mockArticleRepo := new(mocks.ArticleRepository)
	var mockArticle models.Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)

	mockArticleRepo.On("GetByID", mock.AnythingOfType("int64")).Return(&mockArticle, nil)
	defer mockArticleRepo.AssertCalled(t, "GetByID", mock.AnythingOfType("int64"))
	u := ucase.NewArticleUsecase(mockArticleRepo)

	a, err := u.GetByID(mockArticle.ID)

	assert.NoError(t, err)
	assert.NotNil(t, a)

}

func TestStore(t *testing.T) {
	mockArticleRepo := new(mocks.ArticleRepository)
	var mockArticle models.Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)
	//set to 0 because this is test from Client, and ID is an AutoIncreament
	tempMockArticle := mockArticle
	tempMockArticle.ID = 0

	mockArticleRepo.On("GetByTitle", mock.AnythingOfType("string")).Return(nil, models.NOT_FOUND_ERROR)
	mockArticleRepo.On("Store", mock.AnythingOfType("*article.Article")).Return(mockArticle.ID, nil)
	defer mockArticleRepo.AssertCalled(t, "GetByTitle", mock.AnythingOfType("string"))
	defer mockArticleRepo.AssertCalled(t, "Store", mock.AnythingOfType("*article.Article"))

	u := ucase.NewArticleUsecase(mockArticleRepo)

	a, err := u.Store(&tempMockArticle)

	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, mockArticle.Title, tempMockArticle.Title)

}

func TestDelete(t *testing.T) {
	mockArticleRepo := new(mocks.ArticleRepository)
	var mockArticle models.Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)

	mockArticleRepo.On("GetByID", mock.AnythingOfType("int64")).Return(&mockArticle, models.NOT_FOUND_ERROR)
	defer mockArticleRepo.AssertCalled(t, "GetByID", mock.AnythingOfType("int64"))

	mockArticleRepo.On("Delete", mock.AnythingOfType("int64")).Return(true, nil)
	defer mockArticleRepo.AssertCalled(t, "Delete", mock.AnythingOfType("int64"))

	u := ucase.NewArticleUsecase(mockArticleRepo)

	a, err := u.Delete(mockArticle.ID)

	assert.NoError(t, err)
	assert.True(t, a)

}
