package model

import "time"

type Quote struct {
	Eur float64
	Usd float64
	Mxn float64

	RefreshTime time.Time
}
