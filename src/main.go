package main

import (
	"context"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/http/handlers/jwt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/application/usecase/shortener"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/helpers"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/database/repositories/mongodb/miniurls"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/http/handlers/health"
	shortenerHandler "github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/http/handlers/shortener"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/shared"
)

const closeAppTimeout = time.Second * 10

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

	health.NewHealthHandler(e)
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
	miniURLsRepository := miniurls.NewMiniURLsRepository(mongodbCollection)
	shortenerUseCase := shortener.NewShortenerUseCase(miniURLsRepository)
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
