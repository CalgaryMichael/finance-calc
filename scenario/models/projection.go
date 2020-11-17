package scenario

import (
	debtModels "financeCalc/debt/models"
	savingsModels "financeCalc/savings/models"
	"time"
)

type Projection struct {
	EffectiveDate      time.Time
	DebtProjections    []*debtModels.DebtProjection
	SavingsProjections []*savingsModels.SavingsProjection
}

func (projection Projection) OutstandingDebt() bool {
	for _, debtProjection := range projection.DebtProjections {
		if debtProjection.DebtTotal > 0 {
			return true
		}
	}
	return false
}
