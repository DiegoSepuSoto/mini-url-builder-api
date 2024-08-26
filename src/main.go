package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/application/usecase/shortener"
	_ "github.com/DiegoSepuSoto/mini-url-builder-api/src/docs"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/helpers"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/database/repositories/mongodb/miniurls"
	cacheminiurls "github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/database/repositories/redis/miniurls"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/http/handlers/docs"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/http/handlers/health"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/http/handlers/jwt"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/http/handlers/metrics"
	shortenerHandler "github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/http/handlers/shortener"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/shared"
)

const closeAppTimeout = time.Second * 10

// @title Mini URL Builder API
// @version 0.1
// @description This service will create a mini URL and send as a response

// @contact.name Diego Sepúlveda
// @contact.url https://github.com/DiegoSepuSoto
// @contact.email diegosepu.soto@gmail.com

// @host localhost:8080
// @BasePath /
func main() {
	viper.Set("COUNTER_VALUE", -1)
	viper.Set("COUNTER_MAX", -1)

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := helpers.LoadIDRanges()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	metrics.NewMetricsHandler(e)
	health.NewHealthHandler(e)
	docs.NewSwaggerHandler(e)
	jwt.NewJWTHandler(e)

	initShortenerHandler(e)

	quit := make(chan os.Signal, 1)
	go startServer(e, quit)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	gracefulShutdown(e)
}

func initShortenerHandler(e *echo.Echo) {
	mongodbCollection := shared.CreateMongoDBCollection()
	redisClient := shared.CreateRedisClient()

	miniURLsCacheRepository := cacheminiurls.NewMiniURLsCacheRepository(redisClient)
	miniURLsRepository := miniurls.NewMiniURLsRepository(mongodbCollection)
	shortenerUseCase := shortener.NewShortenerUseCase(miniURLsRepository, miniURLsCacheRepository)
	_ = shortenerHandler.NewShortenerHandler(e, shortenerUseCase)
}

func startServer(e *echo.Echo, quit chan os.Signal) {
	log.Print("starting server")

	if err := e.Start(":" + os.Getenv("APP_PORT")); err != nil {
		log.Error(err.Error())
		close(quit)
	}
}

func gracefulShutdown(e *echo.Echo) {
	log.Print("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), closeAppTimeout)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
