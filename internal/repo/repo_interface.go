package repo

import (
	"context"
	"curr-quote/internal/model"
	"github.com/jackc/pgx/v5"
)

type Repo interface {
	GetQuote(ctx context.Context, id string, curr model.Currency) (model.Quote, error)
	SetQuote(ctx context.Context, id string, curr model.Currency, quote model.Quote) error
}

func New(conn *pgx.Conn) Repo {
	return &repoImpl{
		Conn: conn,
	}
}
