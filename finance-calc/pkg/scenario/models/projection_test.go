package scenario

import (
	debtModels "financeCalc/pkg/debt/models"
	savingsModels "financeCalc/pkg/savings/models"
	"testing"
	"time"
)

func Test_OutstandingDebt__WithOutstandingBalance(t *testing.T) {
	projection := Projection{
		EffectiveDate: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		DebtProjections: []*debtModels.DebtProjection{
			&debtModels.DebtProjection{
				Debt:       nil,
				DebtTotal:  50.00,
				PaymentSum: 0.00,
			},
		},
		SavingsProjections: []*savingsModels.SavingsProjection{},
	}
	actual := projection.OutstandingDebt()

	if actual != true {
		t.Error("Projection does not have outstanding debt; expected to have outstanding debt")
	}
}

func Test_OutstandingDebt__NoOutstandingBalance(t *testing.T) {
	projection := Projection{
		EffectiveDate: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		DebtProjections: []*debtModels.DebtProjection{
			&debtModels.DebtProjection{
				Debt:       nil,
				DebtTotal:  0.00,
				PaymentSum: 0.00,
			},
		},
		SavingsProjections: []*savingsModels.SavingsProjection{},
	}
	actual := projection.OutstandingDebt()

	if actual != false {
		t.Error("Projection has outstanding debt; expected to not have outstanding debt")
	}
}
