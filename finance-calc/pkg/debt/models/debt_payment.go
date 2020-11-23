package debt

import "time"

type DebtPayment struct {
	Amount    float64    `json:"amount"`
	CarryOver bool       `json:"carryOver"`
	StartDate *time.Time `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`
}

func (debtPayment DebtPayment) IsActive(currentDate time.Time) bool {
	return (debtPayment.StartDate == nil || currentDate.After(*debtPayment.StartDate)) &&
		(debtPayment.EndDate == nil || currentDate.Before(*debtPayment.EndDate))
}
