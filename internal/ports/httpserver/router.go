package httpserver

import (
	_ "curr-quote/docs"
	"curr-quote/internal/app"
	"github.com/gin-gonic/gin"
)

// routes устанавливает обработчики на эндпоинты
func routes(r *gin.RouterGroup, a app.App) {
	r.PATCH("/quotes/", handleRefreshQuote(a))
	r.GET("/quotes/:id", handleGetQuoteById(a))
	r.GET("/quotes/", handleGetLastQuote(a))
}
