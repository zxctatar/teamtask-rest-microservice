package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"userservice/internal/config"
	bcrypthash "userservice/internal/infrastructure/bcrypt"
	"userservice/internal/infrastructure/postgres"
	"userservice/internal/transport/rest"
	resthandler "userservice/internal/transport/rest/handler"
	"userservice/internal/transport/rest/middleware"
	"userservice/internal/usecase/implementations/registration"
	"userservice/pkg/logger"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// LOAD CONFIG
	config := config.MustLoad()

	// SETUP LOGGER
	log := logger.SetupLogger(config.LogConf.Level)

	// OPEN SQL CONNECTION
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.PostgresConf.Host,
		config.PostgresConf.Port,
		config.PostgresConf.User,
		config.PostgresConf.Password,
		config.PostgresConf.DbName,
		config.PostgresConf.Sslmode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("failed to open database: " + err.Error())
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic("failed to connect to the database: " + err.Error())
	}

	// CREATE POSTGRES
	pos := postgres.NewPostgres(db)
	// CREATE HASHER
	passHasher := bcrypthash.NewBcryptHasher()

	// CREATE REGISTRATION USECASE
	regUC := registration.NewRegUC(log, pos, passHasher)

	// CREATE HANDLER
	handl := resthandler.NewRestHandler(log, regUC)

	// GIN SETTINGS
	gin.SetMode(config.RestConf.Mode)
	router := gin.New()
	router.Use(middleware.TimeoutMiddleware(config.RestConf.RequestTimeout))
	router.Use(gin.Recovery())

	// REGISTER HTTP ROUTES
	router.POST("/registration", handl.Registration)
	router.POST("/login", handl.Login)

	// SERVER SETTING
	serv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.RestConf.Port),
		Handler:      router,
		WriteTimeout: config.RestConf.WriteTimeout,
		ReadTimeout:  config.RestConf.ReadTimeout,
	}

	restServer := rest.NewRestServer(log, serv)

	go restServer.MustStart()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)

	<-quitCh

	ctx, cancel := context.WithTimeout(context.Background(), config.RestConf.ShutdownTimeout*time.Second)
	defer cancel()

	restServer.Stop(ctx)
}
