package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	userHandler "github.com/alifahsanilsatria/twitter-clone/user/delivery/http"
	userMiddleware "github.com/alifahsanilsatria/twitter-clone/user/delivery/http/middleware"
	userRepository "github.com/alifahsanilsatria/twitter-clone/user/repository/db"
	userUsecase "github.com/alifahsanilsatria/twitter-clone/user/usecase"
	"github.com/sirupsen/logrus"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	logger := logrus.StandardLogger()

	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

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
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()

	userMiddleWare := userMiddleware.InitMiddleware()
	e.Use(userMiddleWare.CORS)

	userRepository := userRepository.NewUserRepository(dbConn, logger)
	userUsecase := userUsecase.NewUserUsecase(userRepository, logger)
	userHandler.NewUserHandler(e, userUsecase, logger)

	serverListener := http.Server{
		Addr:    viper.GetString("server.address"),
		Handler: e,
	}

	if err := serverListener.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

	return
}
