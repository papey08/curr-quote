package model

import "time"

type Quote struct {
	/*Eur float64
	Usd float64
	Mxn float64*/
	Data map[Currency]float64

	RefreshTime time.Time
}

func NewQuote() Quote {
	return Quote{
		Data: make(map[Currency]float64),
	}
}
