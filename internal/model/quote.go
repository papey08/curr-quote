package model

import "time"

// Quote содержит все котировки для одной валюты и время обновления
type Quote struct {
	Data map[Currency]float64

	RefreshTime time.Time
}

func NewQuote() Quote {
	return Quote{
		Data: make(map[Currency]float64),
	}
}
