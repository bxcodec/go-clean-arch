package usecase_test

import (
	"context"
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/bxcodec/go-clean-arch/v2/article/mocks"
	ucase "github.com/bxcodec/go-clean-arch/v2/article/usecase"
	_authorMock "github.com/bxcodec/go-clean-arch/v2/author/mocks"
	models "github.com/bxcodec/go-clean-arch/v2/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	mockArticleRepo := new(mocks.ArticleRepository)
	mockArticle := &models.Article{
		Title:   "Hello",
		Content: "Content",
	}

	mockListArtilce := make([]*models.Article, 0)
	mockListArtilce = append(mockListArtilce, mockArticle)
	mockArticleRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(mockListArtilce, nil)
	mockAuthor := &models.Author{
		ID:   1,
		Name: "Iman Tumorang",
	}
	mockAuthorrepo := new(_authorMock.AuthorRepository)
	mockAuthorrepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockAuthor, nil)
	u := ucase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)
	num := int64(1)
	cursor := "12"
	list, nextCursor, err := u.Fetch(context.TODO(), cursor, num)
	cursorExpected := strconv.Itoa(int(mockArticle.ID))
	assert.Equal(t, cursorExpected, nextCursor)
	assert.NotEmpty(t, nextCursor)
	assert.NoError(t, err)
	assert.Len(t, list, len(mockListArtilce))

	mockArticleRepo.AssertExpectations(t)
	mockAuthorrepo.AssertExpectations(t)
}

func TestFetchError(t *testing.T) {
	mockArticleRepo := new(mocks.ArticleRepository)

	mockArticleRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpexted Error"))

	mockAuthorrepo := new(_authorMock.AuthorRepository)
	u := ucase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)
	num := int64(1)
	cursor := "12"
	list, nextCursor, err := u.Fetch(context.TODO(), cursor, num)

	assert.Empty(t, nextCursor)
	assert.Error(t, err)
	assert.Len(t, list, 0)
	mockArticleRepo.AssertExpectations(t)
	mockAuthorrepo.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	mockArticleRepo := new(mocks.ArticleRepository)
	mockArticle := models.Article{
		Title:   "Hello",
		Content: "Content",
	}

	mockArticleRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockArticle, nil)

	mockAuthor := &models.Author{
		ID:   1,
		Name: "Iman Tumorang",
	}
	mockAuthorrepo := new(_authorMock.AuthorRepository)
	mockAuthorrepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockAuthor, nil)
	u := ucase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)

	a, err := u.GetByID(context.TODO(), mockArticle.ID)

	assert.NoError(t, err)
	assert.NotNil(t, a)

	mockArticleRepo.AssertExpectations(t)
	mockAuthorrepo.AssertExpectations(t)
}

func TestStore(t *testing.T) {
	mockArticleRepo := new(mocks.ArticleRepository)
	mockArticle := models.Article{
		Title:   "Hello",
		Content: "Content",
	}
	//set to 0 because this is test from Client, and ID is an AutoIncreament
	tempMockArticle := mockArticle
	tempMockArticle.ID = 0

	mockArticleRepo.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(nil, models.NOT_FOUND_ERROR)
	mockArticleRepo.On("Store", mock.Anything, mock.AnythingOfType("*models.Article")).Return(mockArticle.ID, nil)

	mockAuthorrepo := new(_authorMock.AuthorRepository)
	u := ucase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)

	a, err := u.Store(context.TODO(), &tempMockArticle)

	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, mockArticle.Title, tempMockArticle.Title)
	mockArticleRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockArticleRepo := new(mocks.ArticleRepository)
	mockArticle := models.Article{
		Title:   "Hello",
		Content: "Content",
	}

	mockArticleRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockArticle, models.NOT_FOUND_ERROR)

	mockArticleRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(true, nil)

	mockAuthorrepo := new(_authorMock.AuthorRepository)
	u := ucase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)

	a, err := u.Delete(context.TODO(), mockArticle.ID)

	assert.NoError(t, err)
	assert.True(t, a)
	mockArticleRepo.AssertExpectations(t)
	mockAuthorrepo.AssertExpectations(t)

}
