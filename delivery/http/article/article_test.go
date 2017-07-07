package article_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	articleHttp "github.com/bxcodec/go-clean-arch/delivery/http/article"
	"github.com/bxcodec/go-clean-arch/models"
	"github.com/bxcodec/go-clean-arch/usecase/mocks"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/bxcodec/faker"
	httpHelper "github.com/bxcodec/go-clean-arch/delivery/helper"
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
	handler := articleHttp.ArticleHandler{AUsecase: mockUCase, Helper: httpHelper.HttpHelper{}}
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
	mockUCase.On("Fetch", cursor, int64(num)).Return(nil, "", models.NewErrorInternalServer())

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/article?num=1&cursor="+cursor, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := articleHttp.ArticleHandler{AUsecase: mockUCase, Helper: httpHelper.HttpHelper{}}
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
	handler := articleHttp.ArticleHandler{AUsecase: mockUCase, Helper: httpHelper.HttpHelper{}}
	handler.GetByID(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertCalled(t, "GetByID", int64(num))
}

func TestStore(t *testing.T) {
	mockArticle := models.Article{
		Title:     "Title",
		Content:   "Content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tempMockArticle := mockArticle
	tempMockArticle.ID = 0
	mockUCase := new(mocks.ArticleUsecase)

	j, err := json.Marshal(tempMockArticle)
	assert.NoError(t, err)

	mockUCase.On("Store", &tempMockArticle).Return(&mockArticle, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/article", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/article")

	handler := articleHttp.ArticleHandler{AUsecase: mockUCase, Helper: httpHelper.HttpHelper{}}
	handler.Store(c)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertCalled(t, "Store", &tempMockArticle)
}

func TestDelete(t *testing.T) {
	var mockArticle models.Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)

	mockUCase := new(mocks.ArticleUsecase)

	num := int(mockArticle.ID)

	mockUCase.On("Delete", int64(num)).Return(true, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/article/"+strconv.Itoa(int(num)), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("article/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := articleHttp.ArticleHandler{AUsecase: mockUCase, Helper: httpHelper.HttpHelper{}}
	handler.Delete(c)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUCase.AssertCalled(t, "Delete", int64(num))

}
