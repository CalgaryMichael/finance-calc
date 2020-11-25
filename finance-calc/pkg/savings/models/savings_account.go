package savings

import "time"

type SavingsAccount struct {
	Name           string            `json:"name"`
	APY            float64           `json:"apy"`
	InitialCapital float64           `json:"initialCapital"`
	Payments       []*SavingsPayment `json:"payments"`
	ProjectedDate  *time.Time        `json:"projectedDate"`
}

func (account SavingsAccount) SumActivePayments(currentDate time.Time) float64 {
	totalAmount := 0.0
	for _, payment := range account.Payments {
		if payment.IsActive(currentDate) {
			totalAmount += payment.Amount
		}
	}
	return totalAmount
}
