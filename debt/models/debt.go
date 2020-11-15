package debt

import "time"

type Debt struct {
	DebtName     string
	DebtTotal    float64
	Payments     []*DebtPayment
	InterestRate float64
}

func (debt Debt) IsSettled() bool {
	return debt.DebtTotal == 0
}

func (debt Debt) SumActivePayments(currentDate time.Time) float64 {
	totalAmount := 0.0
	for _, payment := range debt.Payments {
		if payment.IsActive(currentDate) && (debt.DebtTotal > 0 || payment.CarryOver == true) {
			totalAmount += payment.Amount
		}
	}
	return totalAmount
}

func (debt Debt) SumPayments() float64 {
	totalAmount := 0.0
	for _, payment := range debt.Payments {
		totalAmount += payment.Amount
	}
	return totalAmount
}
