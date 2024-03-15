package repo

import (
	"context"
	"curr-quote/internal/model"
)

type Repo interface {
	GetQuote(ctx context.Context, id string, curr model.Currency) (model.Quote, error)
	SetQuote(ctx context.Context, id string, curr model.Currency, quote model.Quote) error
}
