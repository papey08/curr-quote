package repo

import (
	"context"
	"curr-quote/internal/model"
	"errors"
	"github.com/jackc/pgx/v5"
)

var getQuoteQueries = map[model.Currency]string{
	model.EUR: `
		SELECT "refresh_time", "usd", "mxn" FROM "eur_quotes"
		WHERE "id" = $1;`,

	model.USD: `
		SELECT "refresh_time", "eur", "mxn" FROM "usd_quotes"
		WHERE "id" = $1;`,

	model.MXN: `
		SELECT "refresh_time", "eur", "usd" FROM "mxn_quotes"
		WHERE "id" = $1;`,
}

func (r *repoImpl) GetQuote(ctx context.Context, id string, curr model.Currency) (model.Quote, error) {
	quote := model.NewQuote()
	var scanArgs [2]float64

	row := r.QueryRow(ctx, getQuoteQueries[curr], id)
	if err := row.Scan(
		&quote.RefreshTime,
		&scanArgs[0],
		&scanArgs[1],
	); errors.Is(err, pgx.ErrNoRows) {
		return model.Quote{}, model.ErrQuoteNotFound
	} else if err != nil {
		return model.Quote{}, errors.Join(model.ErrDatabaseError, err)
	}

	var i int
	for _, c := range model.SupportableCurrencies {
		if c == curr {
			quote.Data[c] = 1
		} else {
			quote.Data[c] = scanArgs[i]
			i++
		}
	}

	return quote, nil
}
