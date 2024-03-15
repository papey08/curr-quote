package httpserver

import (
	"curr-quote/internal/app"
	"curr-quote/internal/model"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func handleRefreshQuote(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		baseCurr, otherCurr, ok := getCurrencies(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		id, err := a.RefreshQuote(c, baseCurr, otherCurr)

		switch {
		case err == nil:
			c.JSON(http.StatusOK, makeIdResponse(id))
		case errors.Is(err, model.ErrInvalidCurr):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidCurr))
		case errors.Is(err, model.ErrApiError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrApiError))
		case errors.Is(err, model.ErrDatabaseError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrDatabaseError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrServiceError))
		}
	}
}

func handleGetQuoteById(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		quoteValue, err := a.GetQuoteById(c, id)

		switch {
		case err == nil:
			c.JSON(http.StatusOK, makeQuoteResponse(quoteValue.Value, quoteValue.RefreshTime))
		case errors.Is(err, model.ErrQuoteNotFound):
			c.JSON(http.StatusNotFound, errorResponse(model.ErrQuoteNotFound))
		case errors.Is(err, model.ErrInvalidId):
			c.JSON(http.StatusBadRequest, errorResponse(model.ErrInvalidId))
		case errors.Is(err, model.ErrInvalidCurr):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidCurr))
		case errors.Is(err, model.ErrApiError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrApiError))
		case errors.Is(err, model.ErrDatabaseError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrDatabaseError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrServiceError))

		}
	}
}

func handleGetLastQuote(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		baseCurr, otherCurr, ok := getCurrencies(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		quoteValue, err := a.GetLastQuote(c, baseCurr, otherCurr)

		switch {
		case err == nil:
			c.JSON(http.StatusOK, makeQuoteResponse(quoteValue.Value, quoteValue.RefreshTime))
		case errors.Is(err, model.ErrInvalidCurr):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidCurr))
		case errors.Is(err, model.ErrApiError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrApiError))
		case errors.Is(err, model.ErrDatabaseError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrDatabaseError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrServiceError))
		}
	}
}

func getCurrencies(c *gin.Context) (model.Currency, model.Currency, bool) {
	code, ok := c.GetQuery("code")
	if !ok {
		return "", "", false
	}
	codes := strings.Split(code, "/")
	if len(codes) < 2 {
		return "", "", false
	}
	baseCurr := model.Currency(codes[0])
	otherCurr := model.Currency(codes[1])
	return baseCurr, otherCurr, true
}
