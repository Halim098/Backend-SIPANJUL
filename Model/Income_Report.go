package Model

type IncomeReport struct {
	CurrentValue int `json:"currentValue"`
	OldValue int `json:"oldValue"`
	Percentage float64 `json:"percentage"`
	IsNegative bool `json:"isNegative"`
}