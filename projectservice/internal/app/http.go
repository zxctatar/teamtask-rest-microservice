package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"projectservice/internal/config"
	"projectservice/internal/transport/rest"
	resthandler "projectservice/internal/transport/rest/handler"
	"projectservice/internal/transport/rest/middleware"

	"github.com/gin-gonic/gin"
)

func mustLoadHttpServer(cfg *config.Config, log *slog.Logger) *rest.RestServer {
	gin.SetMode(cfg.RestConf.Mode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.TimeoutMiddleware(cfg.RestConf.RequestTimeout))
	handl := resthandler.NewHandler(log)

	router.POST("/project/create", handl.CreateProject)
	router.DELETE("/project/delete", handl.RemoveProject)
	router.GET("/project/getall", handl.GetAllProjects)

	serv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.RestConf.Port),
		Handler:      router,
		ReadTimeout:  cfg.RestConf.ReadTimeout,
		WriteTimeout: cfg.RestConf.WriteTimeout,
	}

	return rest.NewRestServer(log, serv)
}
