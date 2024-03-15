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
	if err = json.Unmarshal(body, &data); err != nil {
		return model.Quote{}, errors.Join(model.ErrApiError, err)
	}

	var res model.Quote
	var ok bool
	if res.Eur, ok = data.Values[strings.ToLower(string(model.EUR))]; !ok {
		return model.Quote{}, errors.Join(model.ErrApiError, errors.New("missing value for eur"))
	}
	if res.Usd, ok = data.Values[strings.ToLower(string(model.USD))]; !ok {
		return model.Quote{}, errors.Join(model.ErrApiError, errors.New("missing value for usd"))
	}
	if res.Mxn, ok = data.Values[strings.ToLower(string(model.MXN))]; !ok {
		return model.Quote{}, errors.Join(model.ErrApiError, errors.New("missing value for mxn"))
	}

	// курсы валют в используемой мной api обновляются раз в сутки,
	// поэтому время последнего обновления задаётся здесь
	res.RefreshTime = time.Now().UTC()

	return res, nil
}

type respBody struct {
	Values map[string]float64 `json:"eur,usd,mxn"`
}
