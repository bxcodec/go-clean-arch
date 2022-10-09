package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	articleHttpDelivery "github.com/bxcodec/go-clean-arch/article/delivery/http"
	httpMiddL "github.com/bxcodec/go-clean-arch/article/delivery/http/middleware"
	articleRepo "github.com/bxcodec/go-clean-arch/article/repository/mysql"
	articleUcase "github.com/bxcodec/go-clean-arch/article/usecase"
	authorRepo "github.com/bxcodec/go-clean-arch/author/repository/mysql"
	"github.com/bxcodec/go-clean-arch/domain"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func main() {
	setupConfig()

	articleUsecase := InitRepositoriesAndUsecase()

	e := echo.New()
	e.Use(httpMiddL.CORS)

	articleHttpDelivery.NewArticleHandler(e, articleUsecase)

	log.Fatal(e.Start(viper.GetString("server.address")))
}

func setupConfig() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func InitRepositoriesAndUsecase() domain.ArticleUsecase {
	dbConn := InitMySQL()
	defer func() {
		if err := dbConn.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	authorRepo := authorRepo.NewMysqlAuthorRepository(dbConn)
	ar := articleRepo.NewMysqlArticleRepository(dbConn)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext)

	return au
}

func InitMySQL() *sql.DB {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		log.Fatal(err)
	}
	
	if err = dbConn.Ping(); err != nil {
		log.Fatal(err)
	}
	return dbConn
}
