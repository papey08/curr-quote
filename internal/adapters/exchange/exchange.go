package exchange

import (
	"context"
	"curr-quote/internal/model"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type exchangeImpl struct {
	http.Client
}

const (
	apiUrl = "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/%s.json"
)

func (e *exchangeImpl) GetLatestQuote(ctx context.Context, curr model.Currency) (model.Quote, error) {
	url := fmt.Sprintf(apiUrl, strings.ToLower(string(curr)))
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := e.Do(req)
	if err != nil {
		return model.Quote{}, errors.Join(model.ErrApiError, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.Quote{}, errors.Join(model.ErrApiError, err)
	}

	var data respBody
	switch curr {
	case model.EUR:
		var eurData respBodyEur
		if err = json.Unmarshal(body, &eurData); err != nil {
			return model.Quote{}, errors.Join(model.ErrApiError, err)
		}
		data = respBody(eurData)
	case model.USD:
		var usdData respBodyUsd
		if err = json.Unmarshal(body, &usdData); err != nil {
			return model.Quote{}, errors.Join(model.ErrApiError, err)
		}
		data = respBody(usdData)
	case model.MXN:
		var mxnData respBodyMxn
		if err = json.Unmarshal(body, &mxnData); err != nil {
			return model.Quote{}, errors.Join(model.ErrApiError, err)
		}
		data = respBody(mxnData)
	}

	res := model.NewQuote()
	var ok bool
	for _, c := range model.SupportableCurrencies {
		if res.Data[c], ok = data.Values[strings.ToLower(string(c))]; !ok {
			return model.Quote{}, errors.Join(model.ErrApiError, errors.New("missing values for some currencies"))
		}
	}

	// курсы валют в используемой мной api обновляются раз в сутки,
	// поэтому время последнего обновления задаётся здесь
	res.RefreshTime = time.Now().UTC()

	return res, nil
}

type respBody struct {
	Values map[string]float64
}
