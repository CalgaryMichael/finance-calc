package scenario

import (
	debtModels "financeCalc/pkg/debt/models"
	savingsModels "financeCalc/pkg/savings/models"
	"time"
)

type Projection struct {
	Id                 int                                `json:"id"`
	EffectiveDate      time.Time                          `json:"effectiveDate"`
	DebtProjections    []*debtModels.DebtProjection       `json:"debtProjections"`
	SavingsProjections []*savingsModels.SavingsProjection `json:"savingsProjections"`
}

func (projection Projection) OutstandingDebt() bool {
	for _, debtProjection := range projection.DebtProjections {
		if debtProjection.DebtTotal > 0 {
			return true
		}
	}
	return false
}
