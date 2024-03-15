package repo

import (
	"context"
	"curr-quote/internal/model"
	"errors"
)

var setQuoteQueries = map[model.Currency]string{
	model.EUR: `
		INSERT INTO "eur_quotes" ("id", "usd", "mxn", "refresh_time") 
		VALUES ($1, $3, $4, $2);`,

	model.USD: `
		INSERT INTO "usd_quotes" ("id", "eur", "mxn", "refresh_time") 
		VALUES ($1, $3, $4, $2);`,

	model.MXN: `
		INSERT INTO "mxn_quotes" ("id", "eur", "usd", "refresh_time") 
		VALUES ($1, $3, $4, $2);`,
}

func (r *repoImpl) SetQuote(ctx context.Context, id string, curr model.Currency, quote model.Quote) error {
	query := setQuoteQueries[curr]
	args := make([]float64, 0, len(model.SupportableCurrencies)-1)

	for _, c := range model.SupportableCurrencies {
		if c != curr {
			args = append(args, quote.Data[c])
		}
	}

	if _, err := r.Exec(ctx, query, id, quote.RefreshTime, args[0], args[1]); err != nil {
		return errors.Join(model.ErrDatabaseError, err)
	}
	return nil
}
