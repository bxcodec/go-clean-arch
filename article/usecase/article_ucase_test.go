package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bxcodec/go-clean-arch/article/mocks"
	ucase "github.com/bxcodec/go-clean-arch/article/usecase"
	_authorMock "github.com/bxcodec/go-clean-arch/author/mocks"
	"github.com/bxcodec/go-clean-arch/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	mockArticleRepo := new(mocks.Repository)
	mockArticle := &models.Article{
		Title:   "Hello",
		Content: "Content",
	}

	mockListArtilce := make([]*models.Article, 0)
	mockListArtilce = append(mockListArtilce, mockArticle)

	t.Run("success", func(t *testing.T) {
		mockArticleRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int64")).Return(mockListArtilce, "next-cursor", nil).Once()
		mockAuthor := &models.Author{
			ID:   1,
			Name: "Iman Tumorang",
		}
		mockAuthorrepo := new(_authorMock.Repository)
		mockAuthorrepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockAuthor, nil)
		u := ucase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)
		num := int64(1)
		cursor := "12"
		list, nextCursor, err := u.Fetch(context.TODO(), cursor, num)
		cursorExpected := "next-cursor"
		assert.Equal(t, cursorExpected, nextCursor)
		assert.NotEmpty(t, nextCursor)
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListArtilce))

		mockArticleRepo.AssertExpectations(t)
		mockAuthorrepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockArticleRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int64")).Return(nil, "", errors.New("Unexpexted Error")).Once()

		mockAuthorrepo := new(_authorMock.Repository)
		u := ucase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)
		num := int64(1)
		cursor := "12"
		list, nextCursor, err := u.Fetch(context.TODO(), cursor, num)

		assert.Empty(t, nextCursor)
		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockArticleRepo.AssertExpectations(t)
		mockAuthorrepo.AssertExpectations(t)
	})

}

func TestGetByID(t *testing.T) {
	mockArticleRepo := new(mocks.Repository)
	mockArticle := models.Article{
		Title:   "Hello",
		Content: "Content",
	}
	mockAuthor := &models.Author{
		ID:   1,
		Name: "Iman Tumorang",
	}

	t.Run("success", func(t *testing.T) {
		mockArticleRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockArticle, nil).Once()
		mockAuthorrepo := new(_authorMock.Repository)
		mockAuthorrepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockAuthor, nil)
		u := ucase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockArticle.ID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockArticleRepo.AssertExpectations(t)
		mockAuthorrepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockArticleRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpected")).Once()

		mockAuthorrepo := new(_authorMock.Repository)
		u := ucase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockArticle.ID)

		assert.Error(t, err)
		assert.Nil(t, a)

		mockArticleRepo.AssertExpectations(t)
		mockAuthorrepo.AssertExpectations(t)
	})

}

func TestStore(t *testing.T) {
	mockArticleRepo := new(mocks.Repository)
	mockArticle := models.Article{
		Title:   "Hello",
		Content: "Content",
	}

	t.Run("success", func(t *testing.T) {
		tempMockArticle := mockArticle
		tempMockArticle.ID = 0
		mockArticleRepo.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(nil, models.ErrNotFound).Once()
		mockArticleRepo.On("Store", mock.Anything, mock.AnythingOfType("*models.Article")).Return(mockArticle.ID, nil).Once()

		mockAuthorrepo := new(_authorMock.Repository)
		u := ucase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)

		a, err := u.Store(context.TODO(), &tempMockArticle)

		assert.NoError(t, err)
		assert.NotNil(t, a)
		assert.Equal(t, mockArticle.Title, tempMockArticle.Title)
		mockArticleRepo.AssertExpectations(t)
	})
	t.Run("existing-title", func(t *testing.T) {
		existingArticle := mockArticle
		mockArticleRepo.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(&existingArticle, nil).Once()
		mockAuthor := &models.Author{
			ID:   1,
			Name: "Iman Tumorang",
		}
		mockAuthorrepo := new(_authorMock.Repository)
		mockAuthorrepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockAuthor, nil)

		u := ucase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)

		a, err := u.Store(context.TODO(), &mockArticle)

		assert.Error(t, err)
		assert.Nil(t, a)
		mockArticleRepo.AssertExpectations(t)
		mockAuthorrepo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	mockArticleRepo := new(mocks.Repository)
	mockArticle := models.Article{
		Title:   "Hello",
		Content: "Content",
	}

	t.Run("success", func(t *testing.T) {
		mockArticleRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockArticle, nil).Once()

		mockArticleRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()

		mockAuthorrepo := new(_authorMock.Repository)
		u := ucase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)

		err := u.Delete(context.TODO(), mockArticle.ID)

		assert.NoError(t, err)
		mockArticleRepo.AssertExpectations(t)
		mockAuthorrepo.AssertExpectations(t)
	})
	t.Run("article-is-not-exist", func(t *testing.T) {
		mockArticleRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, nil).Once()

		mockAuthorrepo := new(_authorMock.Repository)
		u := ucase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)

		err := u.Delete(context.TODO(), mockArticle.ID)

		assert.Error(t, err)
		mockArticleRepo.AssertExpectations(t)
		mockAuthorrepo.AssertExpectations(t)
	})
	t.Run("error-happens-in-db", func(t *testing.T) {
		mockArticleRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpected Error")).Once()

		mockAuthorrepo := new(_authorMock.Repository)
		u := ucase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)

		err := u.Delete(context.TODO(), mockArticle.ID)

		assert.Error(t, err)
		mockArticleRepo.AssertExpectations(t)
		mockAuthorrepo.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	mockArticleRepo := new(mocks.Repository)
	mockArticle := models.Article{
		Title:   "Hello",
		Content: "Content",
		ID:      23,
	}

	t.Run("success", func(t *testing.T) {
		mockArticleRepo.On("Update", mock.Anything, &mockArticle).Once().Return(nil)

		mockAuthorrepo := new(_authorMock.Repository)
		u := ucase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)

		err := u.Update(context.TODO(), &mockArticle)
		assert.NoError(t, err)
		mockArticleRepo.AssertExpectations(t)
	})
}
