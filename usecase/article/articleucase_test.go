package article_test

import (
	"testing"
	"time"

	"github.com/bxcodec/go-clean-arch/models"
	"github.com/bxcodec/go-clean-arch/repository/mocks"
	ucase "github.com/bxcodec/go-clean-arch/usecase/article"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockArticle = &models.Article{
		ID:        int64(10),
		Title:     "Cinta Buta",
		Content:   "Jatuh Cinta Membunuhku",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockCategory = &models.Category{
		ID:        int64(2),
		Name:      "Kehidupan",
		Tag:       "life",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
)

func TestFetch(t *testing.T) {
	mockArticleRepo := new(mocks.ArticleRepository)

	mockListCategory := make([]*models.Category, 0)
	mockListCategory = append(mockListCategory, mockCategory)
	mockListArtilce := make([]*models.Article, 0)
	mockListArtilce = append(mockListArtilce, mockArticle)
	mockArticleRepo.On("Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(mockListArtilce, nil)
	u := ucase.NewArticleUsecase(mockArticleRepo)
	num := int64(1)
	cursor := "12"
	list, nextCursor, err := u.Fetch(cursor, num)
	assert.Equal(t, "10", nextCursor)
	assert.NotEmpty(t, nextCursor)
	assert.NoError(t, err)
	assert.Len(t, list, len(mockListArtilce))

	mockArticleRepo.AssertCalled(t, "Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64"))

}
