package category

import (
	"net/http"
	"strconv"

	httpHelper "github.com/bxcodec/go-clean-arch/delivery/helper"
	"github.com/bxcodec/go-clean-arch/models"
	categoryUcase "github.com/bxcodec/go-clean-arch/usecase"
	"github.com/labstack/echo"
)

type CategoryHandler struct {
	AUsecase categoryUcase.CategoryUsecase
	Helper   httpHelper.HttpHelper
}

func (a *CategoryHandler) FetchCategory(c echo.Context) error {

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

func (a *CategoryHandler) GetByID(c echo.Context) error {

	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	art, err := a.AUsecase.GetByID(id)
	statusCode := a.Helper.GetStatusCode(err)

	if err != nil {
		return c.JSON(statusCode, err.Error())
	}
	return c.JSON(http.StatusOK, art)
}

func (a *CategoryHandler) Store(c echo.Context) error {
	var category models.Category
	err := c.Bind(&category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ar, err := a.AUsecase.Store(&category)
	statusCode := a.Helper.GetStatusCode(err)
	if err != nil {
		return c.JSON(statusCode, err.Error())
	}
	return c.JSON(http.StatusCreated, ar)
}
func (a *CategoryHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	_, err = a.AUsecase.Delete(id)
	statusCode := a.Helper.GetStatusCode(err)
	if err != nil {
		return c.JSON(statusCode, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
