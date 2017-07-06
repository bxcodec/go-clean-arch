package article_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	articleHttp "github.com/bxcodec/go-clean-arch/delivery/http/article"
	"github.com/bxcodec/go-clean-arch/models"
	"github.com/bxcodec/go-clean-arch/usecase/mocks"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/bxcodec/faker"
)

func TestFetch(t *testing.T) {
	var mockArticle models.Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)
	mockUCase := new(mocks.ArticleUsecase)
	mockListArticle := make([]*models.Article, 0)
	mockListArticle = append(mockListArticle, &mockArticle)
	num := 1
	cursor := "2"
	mockUCase.On("Fetch", cursor, int64(num)).Return(mockListArticle, "10", nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/article?num=1&cursor="+cursor, strings.NewReader(""))
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

func TestFetchError(t *testing.T) {
	mockUCase := new(mocks.ArticleUsecase)
	num := 1
	cursor := "2"
	mockUCase.On("Fetch", cursor, int64(num)).Return(nil, "", errors.New("Internal Server Error "))

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/article?num=1&cursor="+cursor, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := articleHttp.ArticleHandler{mockUCase}
	handler.FetchArticle(c)

	responseCursor := rec.Header().Get("X-Cursor")
	assert.Equal(t, "", responseCursor)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertCalled(t, "Fetch", cursor, int64(num))
}

func TestGetByID(t *testing.T) {
	var mockArticle models.Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)

	mockUCase := new(mocks.ArticleUsecase)

	num := int(mockArticle.ID)

	mockUCase.On("GetByID", int64(num)).Return(&mockArticle, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/article/"+strconv.Itoa(int(num)), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("article/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := articleHttp.ArticleHandler{mockUCase}
	handler.GetByID(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertCalled(t, "GetByID", int64(num))
}
