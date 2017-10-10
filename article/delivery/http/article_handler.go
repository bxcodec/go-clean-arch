package http

import (
	"net/http"
	"strconv"

	models "github.com/bxcodec/go-clean-arch/article"

	articleUcase "github.com/bxcodec/go-clean-arch/article/usecase"
	"github.com/labstack/echo"
)

type HttpArticleHandler struct {
	AUsecase articleUcase.ArticleUsecase
}

func (a *HttpArticleHandler) FetchArticle(c echo.Context) error {

	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)

	cursor := c.QueryParam("cursor")

	listAr, nextCursor, err := a.AUsecase.Fetch(cursor, int64(num))

	statusCode := getStatusCode(err)
	if err != nil {
		return c.JSON(statusCode, err.Error())
	}
	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listAr)
}

func (a *HttpArticleHandler) GetByID(c echo.Context) error {

	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	art, err := a.AUsecase.GetByID(id)
	statusCode := getStatusCode(err)

	if err != nil {
		return c.JSON(statusCode, err.Error())
	}
	return c.JSON(http.StatusOK, art)
}

func (a *HttpArticleHandler) Store(c echo.Context) error {
	var article models.Article
	err := c.Bind(&article)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ar, err := a.AUsecase.Store(&article)
	statusCode := getStatusCode(err)
	if err != nil {
		return c.JSON(statusCode, err.Error())
	}
	return c.JSON(http.StatusCreated, ar)
}
func (a *HttpArticleHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	_, err = a.AUsecase.Delete(id)
	statusCode := getStatusCode(err)
	if err != nil {
		return c.JSON(statusCode, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func getStatusCode(err error) int {
	switch err {
	case models.INTERNAL_SERVER_ERROR:
		return http.StatusInternalServerError
	case models.NOT_FOUND_ERROR:
		return http.StatusNotFound
	case models.CONFLIT_ERROR:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func NewArticleHttpHandler(e *echo.Echo, us articleUcase.ArticleUsecase) {
	handler := &HttpArticleHandler{
		AUsecase: us,
	}

	e.GET("/article", handler.FetchArticle)
	e.POST("/article", handler.Store)
	e.GET("/article/:id", handler.GetByID)
	e.DELETE("/article/:id", handler.Delete)

}
