package exchange

type respBodyEur struct {
	Values map[string]float64 `json:"eur"`
}

type respBodyUsd struct {
	Values map[string]float64 `json:"usd"`
}

type respBodyMxn struct {
	Values map[string]float64 `json:"mxn"`
}
