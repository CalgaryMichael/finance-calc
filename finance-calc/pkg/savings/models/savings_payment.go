package savings

import "time"

type SavingsPayment struct {
	Id        int        `json:"id"`
	Amount    float64    `json:"amount"`
	StartDate *time.Time `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`
}

func (payment SavingsPayment) IsActive(currentDate time.Time) bool {
	return (payment.StartDate == nil || currentDate.After(*payment.StartDate)) &&
		(payment.EndDate == nil || currentDate.Before(*payment.EndDate))
}
