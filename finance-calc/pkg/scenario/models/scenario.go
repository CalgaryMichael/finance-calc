package scenario

import (
	debtModels "financeCalc/pkg/debt/models"
	savingsModels "financeCalc/pkg/savings/models"
	"time"
)

type Scenario struct {
	StartDate       time.Time                       `json:"startDate"`
	Debts           []*debtModels.Debt              `json:"debts"`
	SavingsAccounts []*savingsModels.SavingsAccount `json:"savings"`
}
