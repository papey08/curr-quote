package app

import (
	"context"
	"curr-quote/internal/adapters/exchange"
	"curr-quote/internal/model"
	"curr-quote/internal/repo"
	"curr-quote/pkg/logger"
	"time"
)

const refreshPeriod = time.Hour

type App interface {
	RefreshQuote(ctx context.Context, baseCurr model.Currency, otherCurr model.Currency) (string, error)
	GetQuoteById(ctx context.Context, id string) (model.QuoteValue, error)
	GetLastQuote(ctx context.Context, baseCurr model.Currency, otherCurr model.Currency) (model.QuoteValue, error)
}

func New(
	ctx context.Context,
	api exchange.Exchange,
	r repo.Repo,
	logs logger.Logger,
) App {
	a := &appImpl{
		api:    api,
		r:      r,
		logs:   logs,
		quotes: make(map[model.Currency]model.Quote),
	}

	go func() {
		for {
			a.quotes[model.EUR], _ = a.api.GetLatestQuote(ctx, model.EUR)
			a.quotes[model.USD], _ = a.api.GetLatestQuote(ctx, model.USD)
			a.quotes[model.MXN], _ = a.api.GetLatestQuote(ctx, model.MXN)
			logs.Info(nil, "update quotes to latest completed")
			time.Sleep(refreshPeriod)
		}
	}()
	return a
}
