package httpserver

import (
	"time"
)

func errorResponse(err error) idResponse {
	if err == nil {
		return idResponse{}
	}
	errStr := err.Error()
	return idResponse{
		Data: nil,
		Err:  &errStr,
	}
}

func makeIdResponse(id string) idResponse {
	data := idData{Id: id}
	return idResponse{
		Data: &data,
		Err:  nil,
	}
}

func makeQuoteResponse(value float64, refreshTime time.Time) quoteResponse {
	data := quoteData{
		Value:       value,
		RefreshTime: refreshTime.Unix(),
	}
	return quoteResponse{
		Data: &data,
		Err:  nil,
	}
}

type idData struct {
	Id string `json:"id"`
}

type idResponse struct {
	Data *idData `json:"data"`
	Err  *string `json:"error"`
}

type quoteData struct {
	Value       float64 `json:"value"`
	RefreshTime int64   `json:"refresh_time"`
}

type quoteResponse struct {
	Data *quoteData `json:"data"`
	Err  *string    `json:"error"`
}
