package httpserver

import (
	"curr-quote/internal/app"
	"curr-quote/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func New(addr string, a app.App, logs logger.Logger) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	api := router.Group("api/v1")
	api.Use(panicMiddleware(logs))
	api.Use(logMiddleware(logs))
	routes(api, a)

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
