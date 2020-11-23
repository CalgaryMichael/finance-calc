package scenario

import (
	"encoding/json"
	debtModels "financeCalc/pkg/debt/models"
	savingsModels "financeCalc/pkg/savings/models"
	"io"
	"time"
)

type Scenario struct {
	StartDate       time.Time                       `json:"startDate"`
	Debts           []*debtModels.Debt              `json:"debts"`
	SavingsAccounts []*savingsModels.SavingsAccount `json:"savingsAccounts"`
}

func NewFromJSON(v io.ReadCloser) Scenario {
	decoder := json.NewDecoder(v)
	var scenario Scenario
	err := decoder.Decode(&scenario)
	if err != nil {
		panic(err)
	}
	return scenario
}
