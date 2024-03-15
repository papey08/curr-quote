package httpserver

import (
	"curr-quote/internal/model"
	"curr-quote/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func panicMiddleware(logs logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logs.Error(logger.Fields{
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
				}, fmt.Sprintf("panic: %v", r))
				c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrServiceError))
			}
		}()
		c.Next()
	}
}

func logMiddleware(logs logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		logs.Info(logger.Fields{
			"path":   c.Request.URL.Path,
			"method": c.Request.Method,
			"status": c.Writer.Status(),
		}, "")
	}
}
