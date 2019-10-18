package main

import (
	"2019_2_Shtoby_shto/src/config"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts/photo"
	"2019_2_Shtoby_shto/src/dicts/user"
	"2019_2_Shtoby_shto/src/security"
	"context"
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	"github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs" // docs is generated by Swag CLI
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

var (
	initFlag = flag.Bool("initial start", false, "Check your service")
)

var (
	securityService security.HandlerSecurity
	userService     user.HandlerUserService
	photoService    photo.HandlerPhotoService
	dbService       database.InitDBManager
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

func main() {
	flag.Parse()
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	if err := config.InitConfig(); err != nil {
		e.Logger.Error(err)
		os.Exit(1)
	}

	conf := config.GetInstance()

	httpAddr := ":" + strconv.Itoa(conf.Port)
	e.Logger.Info("API Url:", httpAddr)

	dbService = database.Init()
	db, err := dbService.DbConnect("postgres", conf.DbConfig)
	if err != nil {
		e.Logger.Error(err)
		os.Exit(1)
	}
	dm := database.NewDataManager(db)
	defer dm.CloseConnection()

	e.Logger.SetLevel(echoLog.INFO)
	initService(e, dm, conf)
	newServer(e, httpAddr)
	if *initFlag {
		return
	}

	// great shutdown
	go func() {
		switch conf.Port {
		case 443:
			if err := e.StartTLS(httpAddr, "keys/server.crt", "keys/server.key"); err != http.ErrServerClosed {
				e.Logger.Fatalf("HTTPS server ListenAndServe: %v", err)
			}
		default:
			if err := e.Start(httpAddr); err != http.ErrServerClosed {
				e.Logger.Fatalf("HTTP server ListenAndServe: %v", err)
			}
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func newServer(e *echo.Echo, httpAddr string) {
	e.Logger.Info("serving on", httpAddr)

	apiURL := config.GetInstance().FrontendURL
	// securityService.CheckSession,
	e.Use(middleware.Logger(), securityService.CheckSession, middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{apiURL},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Server = &http.Server{
		Addr:           httpAddr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func initService(e *echo.Echo, db database.IDataManager, conf *config.Config) {
	sessionService := security.NewSessionManager(conf.RedisConfig, conf.RedisPass, conf.RedisDbNumber)
	userService = user.CreateInstance(db)
	photoService = photo.CreateInstance(db)
	securityService = security.CreateInstance(sessionService)
	user.NewUserHandler(e, userService, securityService)
	photo.NewPhotoHandler(e, photoService, userService, securityService)
}
