package http

import (
	httpHelper "github.com/bxcodec/go-clean-arch/delivery/helper"
	artHandler "github.com/bxcodec/go-clean-arch/delivery/http/article"
	catHandler "github.com/bxcodec/go-clean-arch/delivery/http/category"
	usecase "github.com/bxcodec/go-clean-arch/usecase"
	"github.com/labstack/echo"
)

type HttpHandler struct {
	E *echo.Echo
}

func (h *HttpHandler) InitArticleDelivery(au usecase.ArticleUsecase) *HttpHandler {
	helper := httpHelper.HttpHelper{}
	articleHandler := &artHandler.ArticleHandler{AUsecase: au, Helper: helper}
	h.E.GET(`/article`, articleHandler.FetchArticle)
	h.E.GET(`/article/:id`, articleHandler.GetByID)
	h.E.POST(`/article`, articleHandler.Store)
	h.E.DELETE(`/article/:id`, articleHandler.Delete)
	return h
}

func (h *HttpHandler) InitCategoryDelivery(au usecase.CategoryUsecase) *HttpHandler {
	helper := httpHelper.HttpHelper{}
	categoryHandler := &catHandler.CategoryHandler{AUsecase: au, Helper: helper}
	h.E.GET(`/category`, categoryHandler.FetchCategory)
	h.E.GET(`/category/:id`, categoryHandler.GetByID)
	h.E.POST(`/category`, categoryHandler.Store)
	h.E.DELETE(`/category/:id`, categoryHandler.Delete)
	return h
}
