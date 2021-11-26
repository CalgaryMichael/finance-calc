package savings

import (
	debt "financeCalc/pkg/debt"
	debtModels "financeCalc/pkg/debt/models"
	models "financeCalc/pkg/savings/models"
	"time"
)

func ProjectSavingsForMonth(
	accounts []*models.SavingsAccount,
	debtProjections []*debtModels.DebtProjection,
	currentDate time.Time,
) []*models.SavingsProjection {
	projections := make([]*models.SavingsProjection, len(accounts))
	for i, account := range accounts {
		carryOverSum := 0.00
		if i == 0 {
			carryOverSum = debt.GetCarryOverSum(debtProjections)
		}
		projections[i] = buildProjection(account, currentDate, carryOverSum)
	}
	return projections
}

func buildProjection(account *models.SavingsAccount, currentDate time.Time, carryOverSum float64) *models.SavingsProjection {
	payment := sumPayment(account, currentDate, carryOverSum)
	return &models.SavingsProjection{
		SavingsAccount: account,
		SavingsTotal:   account.InitialCapital + payment,
		PaymentSum:     payment,
	}
}

func sumPayment(account *models.SavingsAccount, currentDate time.Time, carryOverSum float64) float64 {
	interest := account.InitialCapital * (account.APY / 12)
	return account.SumActivePayments(currentDate) + carryOverSum + interest
}

func OutstandingSavingsProjections(savingsAccounts []*models.SavingsAccount, currentDate time.Time) bool {
	for _, savingsAccount := range savingsAccounts {
		if !savingsAccount.SatisfiesProjection(currentDate) {
			return true
		}
	}
	return false
}
