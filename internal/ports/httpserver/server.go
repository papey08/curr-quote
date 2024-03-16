package httpserver

import (
	"curr-quote/internal/app"
	"curr-quote/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// New создаёт экземпляр http-сервера с middlewares и эндпоинтами
func New(addr string, a app.App, logs logger.Logger) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("api/v1")
	api.Use(panicMiddleware(logs))
	api.Use(logMiddleware(logs))
	routes(api, a)

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
