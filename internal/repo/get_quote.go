package repo

import (
	"context"
	"curr-quote/internal/model"
	"errors"
	"github.com/jackc/pgx/v5"
)

const (
	getEurQuoteQuery = `
		SELECT "refresh_time", "usd", "mxn" FROM "eur_quotes"
		WHERE "id" = $1;`

	getUsdQuoteQuery = `
		SELECT "refresh_time", "eur", "mxn" FROM "usd_quotes"
		WHERE "id" = $1;`

	getMxnQuoteQuery = `
		SELECT "refresh_time", "eur", "usd" FROM "mxn_quotes"
		WHERE "id" = $1;`
)

func (r *repoImpl) GetQuote(ctx context.Context, id string, curr model.Currency) (model.Quote, error) {
	var quote model.Quote
	var query string
	var scanArgs [2]*float64
	switch curr {
	case model.EUR:
		query = getEurQuoteQuery
		scanArgs = [2]*float64{&quote.Usd, &quote.Mxn}
	case model.USD:
		query = getUsdQuoteQuery
		scanArgs = [2]*float64{&quote.Eur, &quote.Mxn}
	case model.MXN:
		query = getMxnQuoteQuery
		scanArgs = [2]*float64{&quote.Eur, &quote.Usd}
	}
	row := r.QueryRow(ctx, query, id)
	if err := row.Scan(
		&quote.RefreshTime,
		scanArgs[0], // TODO test unpack
		scanArgs[1],
	); errors.Is(err, pgx.ErrNoRows) {
		return model.Quote{}, model.ErrQuoteNotFound
	} else if err != nil {
		return model.Quote{}, errors.Join(model.ErrDatabaseError, err)
	}
	return quote, nil
}
