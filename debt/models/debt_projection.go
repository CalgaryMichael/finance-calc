package debt

type DebtProjection struct {
	Debt         *Debt
	DebtTotal    float64
	PaymentSum   float64
	UnappliedSum float64
}

func (projection DebtProjection) IsSettled() bool {
	return projection.DebtTotal == 0.00 && (projection.PaymentSum >= projection.Debt.DebtTotal || projection.PaymentSum == 0.00)
}
