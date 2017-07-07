package http

import (
	artHandler "github.com/bxcodec/go-clean-arch/delivery/http/article"
	httpHelper "github.com/bxcodec/go-clean-arch/delivery/http/helper"
	artUcase "github.com/bxcodec/go-clean-arch/usecase"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo, au artUcase.ArticleUsecase) {
	helper := httpHelper.HttpHelper{}
	articleHandler := &artHandler.ArticleHandler{AUsecase: au, Helper: helper}
	e.GET(`/article`, articleHandler.FetchArticle)
	e.GET(`/article/:id`, articleHandler.GetByID)
	e.POST(`/article`, articleHandler.Store)
}
