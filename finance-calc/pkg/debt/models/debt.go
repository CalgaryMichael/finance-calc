package debt

import "time"

type Debt struct {
	Id           int            `json:"id"`
	DebtName     string         `json:"name"`
	DebtTotal    float64        `json:"total"`
	Payments     []*DebtPayment `json:"payments"`
	InterestRate float64        `json:"interestRate"`
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
