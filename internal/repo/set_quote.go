package repo

import (
	"context"
	"curr-quote/internal/model"
	"errors"
)

const (
	setEurQuoteQuery = `
		INSERT INTO "eur_quotes" ("id", "usd", "mxn", "refresh_time") 
		VALUES ($1, $3, $4, $2);`

	setUsdQuoteQuery = `
		INSERT INTO "usd_quotes" ("id", "eur", "mxn", "refresh_time") 
		VALUES ($1, $3, $4, $2);`

	setMxnQuoteQuery = `
		INSERT INTO "mxn_quotes" ("id", "eur", "usd", "refresh_time") 
		VALUES ($1, $3, $4, $2);`
)

func (r *repoImpl) SetQuote(ctx context.Context, id string, curr model.Currency, quote model.Quote) error {
	var query string
	var args [2]float64
	switch curr {
	case model.EUR:
		query = setEurQuoteQuery
		args = [2]float64{quote.Usd, quote.Mxn}
	case model.USD:
		query = setUsdQuoteQuery
		args = [2]float64{quote.Eur, quote.Mxn}
	case model.MXN:
		query = setMxnQuoteQuery
		args = [2]float64{quote.Eur, quote.Usd}
	}

	// TODO test unpack
	if _, err := r.Exec(ctx, query, id, quote.RefreshTime, args[0], args[1]); err != nil {
		return errors.Join(model.ErrDatabaseError, err)
	}
	return nil
}
