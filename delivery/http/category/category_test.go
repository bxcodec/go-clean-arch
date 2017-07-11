package category_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	categoryHttp "github.com/bxcodec/go-clean-arch/delivery/http/category"
	"github.com/bxcodec/go-clean-arch/models"
	"github.com/bxcodec/go-clean-arch/usecase/mocks"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/bxcodec/faker"
	httpHelper "github.com/bxcodec/go-clean-arch/delivery/helper"
)

func TestFetch(t *testing.T) {
	var mockCategory models.Category
	err := faker.FakeData(&mockCategory)
	assert.NoError(t, err)
	mockUCase := new(mocks.CategoryUsecase)
	mockListCategory := make([]*models.Category, 0)
	mockListCategory = append(mockListCategory, &mockCategory)
	num := 1
	cursor := "2"
	mockUCase.On("Fetch", cursor, int64(num)).Return(mockListCategory, "10", nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/category?num=1&cursor="+cursor, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := categoryHttp.CategoryHandler{AUsecase: mockUCase, Helper: httpHelper.HttpHelper{}}
	handler.FetchCategory(c)

	responseCursor := rec.Header().Get("X-Cursor")
	assert.Equal(t, "10", responseCursor)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertCalled(t, "Fetch", cursor, int64(num))
}

func TestFetchError(t *testing.T) {
	mockUCase := new(mocks.CategoryUsecase)
	num := 1
	cursor := "2"
	mockUCase.On("Fetch", cursor, int64(num)).Return(nil, "", models.NewErrorInternalServer())

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/category?num=1&cursor="+cursor, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := categoryHttp.CategoryHandler{AUsecase: mockUCase, Helper: httpHelper.HttpHelper{}}
	handler.FetchCategory(c)

	responseCursor := rec.Header().Get("X-Cursor")
	assert.Equal(t, "", responseCursor)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertCalled(t, "Fetch", cursor, int64(num))
}

func TestGetByID(t *testing.T) {
	var mockCategory models.Category
	err := faker.FakeData(&mockCategory)
	assert.NoError(t, err)

	mockUCase := new(mocks.CategoryUsecase)

	num := int(mockCategory.ID)

	mockUCase.On("GetByID", int64(num)).Return(&mockCategory, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/category/"+strconv.Itoa(int(num)), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("category/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := categoryHttp.CategoryHandler{AUsecase: mockUCase, Helper: httpHelper.HttpHelper{}}
	handler.GetByID(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertCalled(t, "GetByID", int64(num))
}

func TestStore(t *testing.T) {
	mockCategory := models.Category{
		Name:      "Name",
		Tag:       "Tag",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tempMockCategory := mockCategory
	tempMockCategory.ID = 0
	mockUCase := new(mocks.CategoryUsecase)

	j, err := json.Marshal(tempMockCategory)
	assert.NoError(t, err)

	mockUCase.On("Store", &tempMockCategory).Return(&mockCategory, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/category", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/category")

	handler := categoryHttp.CategoryHandler{AUsecase: mockUCase, Helper: httpHelper.HttpHelper{}}
	handler.Store(c)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertCalled(t, "Store", &tempMockCategory)
}

func TestDelete(t *testing.T) {
	var mockCategory models.Category
	err := faker.FakeData(&mockCategory)
	assert.NoError(t, err)

	mockUCase := new(mocks.CategoryUsecase)

	num := int(mockCategory.ID)

	mockUCase.On("Delete", int64(num)).Return(true, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/category/"+strconv.Itoa(int(num)), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("category/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := categoryHttp.CategoryHandler{AUsecase: mockUCase, Helper: httpHelper.HttpHelper{}}
	handler.Delete(c)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUCase.AssertCalled(t, "Delete", int64(num))

}
