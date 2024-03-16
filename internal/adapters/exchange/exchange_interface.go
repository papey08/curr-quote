package exchange

import (
	"context"
	"curr-quote/internal/model"
	"net/http"
)

// Exchange представляет собой обёртку для API получения свежих котировок
type Exchange interface {
	// GetLatestQuote возвращает котировки для заданной валюты
	GetLatestQuote(ctx context.Context, curr model.Currency) (model.Quote, error)
}

func New() Exchange {
	return &exchangeImpl{
		Client: http.Client{},
	}
}
