package model

type Currency string

const (
	EUR Currency = "EUR"
	USD Currency = "USD"
	MXN Currency = "MXN"
)

var SupportableCurrencies = [3]Currency{EUR, USD, MXN}
