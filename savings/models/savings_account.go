package savings

import "time"

type SavingsAccount struct {
	Name           string
	APY            float64
	InitialCapital float64
	Payments       []*SavingsPayment
	ProjectedDate  time.Time
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
