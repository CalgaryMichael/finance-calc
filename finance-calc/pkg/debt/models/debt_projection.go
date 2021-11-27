package debt

import "time"

type DebtProjection struct {
	Debt          *Debt     `json:"debt"`
	EffectiveDate time.Time `json:"effectiveDate"`
	DebtTotal     float64   `json:"debtTotal"`
	PaymentSum    float64   `json:"paymentSum"`
	UnappliedSum  float64
}

func (projection DebtProjection) IsSettled() bool {
	return projection.DebtTotal == 0.00 && (projection.PaymentSum >= projection.Debt.DebtTotal || projection.PaymentSum == 0.00)
}

func GetDebtProjectionsForDate(projections []*DebtProjection, date time.Time) []*DebtProjection {
	var filtered []*DebtProjection
	for _, projection := range projections {
		if projection.EffectiveDate == date {
			filtered = append(filtered, projection)
		}
	}
	return filtered
}
