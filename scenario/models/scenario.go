package scenario

import (
	debtModels "financeCalc/debt/models"
	savingsModels "financeCalc/savings/models"
	"time"
)

type Scenario struct {
	StartDate       time.Time
	Debts           []*debtModels.Debt
	SavingsAccounts []*savingsModels.SavingsAccount
}
