package article

import (
	"net/http"
	"strconv"

	articleUcase "github.com/bxcodec/go-clean-arch/usecase/article"
	"github.com/labstack/echo"
)

type ArticleHandler struct {
	AUsecase articleUcase.ArticleUsecase
}

func (a *ArticleHandler) FetchArticle(c echo.Context) error {

	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)

	cursor := c.QueryParam("cursor")

	listAr, nextCursor, err := a.AUsecase.Fetch(cursor, int64(num))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something Error On Our Services")
	}
	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listAr)
}
