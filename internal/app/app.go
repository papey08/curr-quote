package app

import (
	"context"
	"curr-quote/internal/adapters/exchange"
	"curr-quote/internal/model"
	"curr-quote/internal/repo"
	"curr-quote/pkg/logger"
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

type appImpl struct {
	api  exchange.Exchange
	r    repo.Repo
	logs logger.Logger

	quotes map[model.Currency]model.Quote
}

func (a *appImpl) RefreshQuote(ctx context.Context, baseCurr model.Currency, otherCurr model.Currency) (string, error) {
	var err error
	defer func() {
		a.writeLog("RefreshQuote", err)
	}()

	if !(validateCurr(baseCurr) && validateCurr(otherCurr)) {
		return "", model.ErrInvalidCurr
	}

	id := a.generateId(baseCurr, otherCurr)

	go func() {
		defer func() {
			a.writeLog("RefreshQuote", err)
		}()

		var quote model.Quote
		quote, err = a.api.GetLatestQuote(ctx, baseCurr)
		if err != nil {
			return
		}

		dbId := strings.Split(id, "-")[2]
		curr := model.Currency(strings.Split(id, "-")[0])
		err = a.r.SetQuote(ctx, dbId, curr, quote)
		if err != nil {
			return
		}

		a.quotes[baseCurr] = quote
	}()
	return id, nil
}

func (a *appImpl) GetQuoteById(ctx context.Context, id string) (model.QuoteValue, error) {
	var err error
	defer func() {
		a.writeLog("GetQuoteById", err)
	}()

	idData := strings.Split(id, "-")
	if len(idData) != 3 {
		return model.QuoteValue{}, model.ErrInvalidId
	}
	baseCurr := model.Currency(idData[0])
	otherCurr := model.Currency(idData[1])
	dbId := idData[2]
	if !(validateCurr(baseCurr) && validateCurr(otherCurr)) {
		return model.QuoteValue{}, model.ErrInvalidCurr
	}

	var quote model.Quote
	quote, err = a.r.GetQuote(ctx, dbId, baseCurr)
	if err != nil {
		return model.QuoteValue{}, err
	}

	var res model.QuoteValue
	res.Value = quote.Data[otherCurr]
	res.RefreshTime = quote.RefreshTime
	return res, nil
}

func (a *appImpl) GetLastQuote(_ context.Context, baseCurr model.Currency, otherCurr model.Currency) (model.QuoteValue, error) {
	var err error
	defer func() {
		a.writeLog("GetLastQuote", err)
	}()
	if !(validateCurr(baseCurr) && validateCurr(otherCurr)) {
		return model.QuoteValue{}, model.ErrInvalidCurr
	}

	quote := a.quotes[baseCurr]
	var res model.QuoteValue
	res.Value = quote.Data[otherCurr]
	res.RefreshTime = quote.RefreshTime
	return res, nil
}

func (a *appImpl) generateId(baseCurr model.Currency, otherCurr model.Currency) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	buffer := make([]byte, 15)
	for i := range buffer {
		buffer[i] = charset[rand.Intn(len(charset))]
	}
	return fmt.Sprintf("%s-%s-%s", string(baseCurr), string(otherCurr), string(buffer))
}

func validateCurr(curr model.Currency) bool {
	switch curr {
	case model.EUR:
	case model.USD:
	case model.MXN:
	default:
		return false
	}
	return true
}

func (a *appImpl) writeLog(method string, err error) {
	if errors.Is(err, model.ErrApiError) || errors.Is(err, model.ErrDatabaseError) {
		a.logs.Error(logger.Fields{
			"method": method,
		}, err.Error())
	} else if err != nil {
		a.logs.Info(logger.Fields{
			"method": method,
		}, err.Error())
	} else {
		a.logs.Info(logger.Fields{
			"method": method,
		}, "ok")
	}
}
