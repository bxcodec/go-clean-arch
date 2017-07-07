package article

import (
	"net/http"
	"strconv"

	httpHelper "github.com/bxcodec/go-clean-arch/delivery/http/helper"
	"github.com/bxcodec/go-clean-arch/models"
	articleUcase "github.com/bxcodec/go-clean-arch/usecase"
	"github.com/labstack/echo"
)

type ArticleHandler struct {
	AUsecase articleUcase.ArticleUsecase
	Helper   httpHelper.HttpHelper
}

func (a *ArticleHandler) FetchArticle(c echo.Context) error {

	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)

	cursor := c.QueryParam("cursor")

	listAr, nextCursor, err := a.AUsecase.Fetch(cursor, int64(num))

	statusCode := a.Helper.GetStatusCode(err)
	if err != nil {
		return c.JSON(statusCode, err.Error())
	}
	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listAr)
}

func (a *ArticleHandler) GetByID(c echo.Context) error {

	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	art, err := a.AUsecase.GetByID(id)
	statusCode := a.Helper.GetStatusCode(err)

	if err != nil {
		return c.JSON(statusCode, err.Error())
	}
	return c.JSON(http.StatusOK, art)
}

func (a *ArticleHandler) Store(c echo.Context) error {
	var article models.Article
	err := c.Bind(&article)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ar, err := a.AUsecase.Store(&article)
	statusCode := a.Helper.GetStatusCode(err)
	if err != nil {
		return c.JSON(statusCode, err.Error())
	}
	return c.JSON(http.StatusCreated, ar)
}
