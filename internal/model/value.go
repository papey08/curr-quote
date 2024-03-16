package model

import "time"

// QuoteValue содержит отношение одной валюты к другой и время обновления
type QuoteValue struct {
	Value       float64
	RefreshTime time.Time
}
