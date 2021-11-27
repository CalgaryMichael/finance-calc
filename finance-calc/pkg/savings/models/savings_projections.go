package savings

import "time"

type SavingsProjection struct {
	SavingsAccount *SavingsAccount `json:"savingsAccount"`
	EffectiveDate  time.Time       `json:"effectiveDate"`
	SavingsTotal   float64         `json:"savingsTotal"`
	PaymentSum     float64         `json:"paymentSum"`
}

func GetSavingsProjectionsForDate(projections []*SavingsProjection, date time.Time) []*SavingsProjection {
	var filtered []*SavingsProjection
	for _, projection := range projections {
		if projection.EffectiveDate == date {
			filtered = append(filtered, projection)
		}
	}
	return filtered
}
