package app

import (
	"context"
	"curr-quote/internal/adapters/exchange"
	"curr-quote/internal/model"
	"curr-quote/internal/repo"
	"curr-quote/pkg/logger"
	"time"
)

// refreshPeriod определяет периодичность автоматического обновления котировок
const refreshPeriod = time.Hour

// App представляет собой слой бизнес-логики приложения
type App interface {
	RefreshQuote(ctx context.Context, baseCurr model.Currency, otherCurr model.Currency) (string, error)
	GetQuoteById(ctx context.Context, id string) (model.QuoteValue, error)
	GetLastQuote(ctx context.Context, baseCurr model.Currency, otherCurr model.Currency) (model.QuoteValue, error)
}

// New создаёт экземпляр приложения
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

	for _, c := range model.SupportableCurrencies {
		a.quotes[c] = model.NewQuote()
	}

	go func() {
		for {
			for _, k := range model.SupportableCurrencies {
				a.quotes[k], _ = a.api.GetLatestQuote(ctx, k)
			}
			logs.Info(nil, "update quotes to latest completed")
			time.Sleep(refreshPeriod)
		}
	}()
	return a
}
