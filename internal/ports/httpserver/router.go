package httpserver

import (
	"curr-quote/internal/app"
	"github.com/gin-gonic/gin"
)

func routes(r *gin.RouterGroup, a app.App) {
	r.PATCH("/quotes/", handleRefreshQuote(a))
	r.GET("/quotes/:id", handleGetQuoteById(a))
	r.GET("/quotes/", handleGetLastQuote(a))
}
