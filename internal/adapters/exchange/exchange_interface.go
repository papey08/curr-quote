package exchange

import (
	"context"
	"curr-quote/internal/model"
)

type Exchange interface {
	GetLatestQuote(ctx context.Context, curr model.Currency) (model.Quote, error)
}
