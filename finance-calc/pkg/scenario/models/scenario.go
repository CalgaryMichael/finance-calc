package scenario

import (
	debtModels "financeCalc/pkg/debt/models"
	savingsModels "financeCalc/pkg/savings/models"
	"time"
)

type Scenario struct {
	Id              int                             `json:"id"`
	StartDate       time.Time                       `json:"startDate"`
	Debts           []*debtModels.Debt              `json:"debts"`
	SavingsAccounts []*savingsModels.SavingsAccount `json:"savings"`
	SortKey         string                          `json:"sortKey"`
	ReverseSort     bool                            `json:"reverseSort"`
}
