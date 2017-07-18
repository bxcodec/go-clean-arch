package main

import (
	"database/sql"
	"fmt"

	cfg "github.com/bxcodec/go-clean-arch/config/env"
	"github.com/bxcodec/go-clean-arch/config/middleware"
	httpDeliver "github.com/bxcodec/go-clean-arch/delivery/http"
	articleRepo "github.com/bxcodec/go-clean-arch/repository/mysql/article"
	categoryRepo "github.com/bxcodec/go-clean-arch/repository/mysql/category"
	articleUcase "github.com/bxcodec/go-clean-arch/usecase/article"
	categoryUcase "github.com/bxcodec/go-clean-arch/usecase/category"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

var config cfg.Config

func init() {
	config = cfg.NewViperConfig()

	if config.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

}

func main() {

	dbHost := config.GetString(`database.host`)
	dbPort := config.GetString(`database.port`)
	dbUser := config.GetString(`database.user`)
	dbPass := config.GetString(`database.pass`)
	dbName := config.GetString(`database.name`)
	dsn := dbUser + `:` + dbPass + `@tcp(` + dbHost + `:` + dbPort + `)/` + dbName + `?parseTime=1&loc=Asia%2FJakarta`
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil && config.GetBool("debug") {
		fmt.Println(err)
	}
	defer dbConn.Close()
	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)
	ar := articleRepo.NewMysqlArticleRepository(dbConn)
	au := articleUcase.NewArticleUsecase(ar)
	cr := categoryRepo.NewMysqlCategoryRepository(dbConn)
	cu := categoryUcase.NewCategoryUsecase(cr)

	httpDeliveryHandler := httpDeliver.HttpHandler{
		E: e,
	}
	httpDeliveryHandler.
		InitArticleDelivery(au).
		InitCategoryDelivery(cu)

	e.Start(config.GetString("server.address"))
}
