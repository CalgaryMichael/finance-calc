package savings

import "time"

type SavingsPayment struct {
	Amount    float64
	StartDate *time.Time
	EndDate   *time.Time
}

func (payment SavingsPayment) IsActive(currentDate time.Time) bool {
	return (payment.StartDate == nil || currentDate.After(*payment.StartDate)) &&
		(payment.EndDate == nil || currentDate.Before(*payment.EndDate))
}
