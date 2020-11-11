package debt

type DebtProjection struct {
	Debt       *Debt
	DebtTotal  float64
	PaymentSum float64
}
