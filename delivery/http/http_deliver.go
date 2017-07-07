package http

import (
	httpHelper "github.com/bxcodec/go-clean-arch/delivery/helper"
	artHandler "github.com/bxcodec/go-clean-arch/delivery/http/article"
	artUcase "github.com/bxcodec/go-clean-arch/usecase"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo, au artUcase.ArticleUsecase) {
	helper := httpHelper.HttpHelper{}
	articleHandler := &artHandler.ArticleHandler{AUsecase: au, Helper: helper}
	e.GET(`/article`, articleHandler.FetchArticle)
	e.GET(`/article/:id`, articleHandler.GetByID)
	e.POST(`/article`, articleHandler.Store)
	e.DELETE(`/article/:id`, articleHandler.Delete)
}
