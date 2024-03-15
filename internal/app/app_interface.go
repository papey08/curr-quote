package app

import (
	"context"
	"curr-quote/internal/model"
)

type App interface {
	RefreshQuote(ctx context.Context, baseCurr model.Currency, otherCurr model.Currency) (string, error)
	GetQuoteById(ctx context.Context, id string) (model.QuoteValue, error)
	GetLastQuote(ctx context.Context, baseCurr model.Currency, otherCurr model.Currency) (model.QuoteValue, error)
}
