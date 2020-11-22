package debt

import "time"

type DebtPayment struct {
	Amount    float64
	CarryOver bool
	StartDate *time.Time
	EndDate   *time.Time
}

func (debtPayment DebtPayment) IsActive(currentDate time.Time) bool {
	return (debtPayment.StartDate == nil || currentDate.After(*debtPayment.StartDate)) &&
		(debtPayment.EndDate == nil || currentDate.Before(*debtPayment.EndDate))
}
