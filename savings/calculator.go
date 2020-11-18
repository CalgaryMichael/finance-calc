package savings

import (
	debt "financeCalc/debt"
	debtModels "financeCalc/debt/models"
	models "financeCalc/savings/models"
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
			carryOverSum = debt.GetCarryOverSum(debtProjections, currentDate)
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
