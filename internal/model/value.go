package model

import "time"

type QuoteValue struct {
	Value       float64
	RefreshTime time.Time
}
