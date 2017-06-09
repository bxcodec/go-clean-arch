package article_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	articleHttp "github.com/bxcodec/go-clean-arch/delivery/http/article"
	"github.com/bxcodec/go-clean-arch/models"
	"github.com/bxcodec/go-clean-arch/usecase/article/mocks"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mockArticle = &models.Article{
		ID:      int64(10),
		Title:   "Cinta Buta",
		Content: "Jatuh Cinta Membunuhku",
	}
)

func TestFetch(t *testing.T) {
	mockUCase := new(mocks.ArticleUsecase)
	mockListArticle := make([]*models.Article, 0)
	mockListArticle = append(mockListArticle, mockArticle)
	num := 1
	cursor := "2"
	mockUCase.On("Fetch", cursor, int64(num)).Return(mockListArticle, "10", nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/article?num="+strconv.Itoa(num)+"&cursor="+cursor, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := articleHttp.ArticleHandler{mockUCase}
	handler.FetchArticle(c)

	responseCursor := rec.Header().Get("X-Cursor")
	assert.Equal(t, "10", responseCursor)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertCalled(t, "Fetch", cursor, int64(num))
}
