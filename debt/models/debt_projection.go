package debt

import (
	"math"
	"time"
)

type DebtProjection struct {
	Debt       *Debt
	DebtTotal  float64
	PaymentSum float64
}

func (projection DebtProjection) IsSettled() bool {
	return projection.DebtTotal == 0.00 && (projection.PaymentSum >= projection.Debt.DebtTotal || projection.PaymentSum == 0.00)
}

// Get the sum of all active payments that are applied to settled debts
func (projection DebtProjection) GetActiveSettledPayments(currentDate time.Time) float64 {
	if projection.IsSettled() {
		paymentSum := projection.Debt.SumActivePayments(currentDate) - projection.Debt.DebtTotal
		return math.Max(paymentSum, 0.00)
	}
	return 0.00
}
