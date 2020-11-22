package scenario

import (
	debtModels "financeCalc/pkg/debt/models"
	savingsModels "financeCalc/pkg/savings/models"
	"time"
)

type Scenario struct {
	StartDate       time.Time
	Debts           []*debtModels.Debt
	SavingsAccounts []*savingsModels.SavingsAccount
}
