package model

type Currency string

const (
	EUR Currency = "EUR"
	USD Currency = "USD"
	MXN Currency = "MXN"
)

// SupportableCurrencies определяет список поддерживаемых валют. Очерёдность
// валют важна для корректной работы с БД
var SupportableCurrencies = [3]Currency{EUR, USD, MXN}
