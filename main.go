package main

import (
	"database/sql"
	"fmt"
	"net/url"

	httpDeliver "github.com/bxcodec/go-clean-arch/article/delivery/http"
	articleRepo "github.com/bxcodec/go-clean-arch/article/repository"
	articleUcase "github.com/bxcodec/go-clean-arch/article/usecase"
	cfg "github.com/bxcodec/go-clean-arch/config/env"
	"github.com/bxcodec/go-clean-arch/config/middleware"
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
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
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

	httpDeliver.NewArticleHttpHandler(e, au)

	e.Start(config.GetString("server.address"))
}
