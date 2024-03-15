package exchange

import (
	"context"
	"curr-quote/internal/model"
	"net/http"
)

type Exchange interface {
	GetLatestQuote(ctx context.Context, curr model.Currency) (model.Quote, error)
}

func New() Exchange {
	return &exchangeImpl{
		Client: http.Client{},
	}
}
