package scenario

import (
	debt "financeCalc/debt"
	debtModels "financeCalc/debt/models"
	savings "financeCalc/savings"
	savingsModels "financeCalc/savings/models"
	models "financeCalc/scenario/models"
	"time"
)

func BuildProjections(scenario models.Scenario, key string, reverse bool) []*models.Projection {
	projections := make([]*models.Projection, 0, 240)

	debts := debt.SortDebts(scenario.Debts, key, reverse)
	savingsAccounts := scenario.SavingsAccounts

	i := 0
	for {
		effectiveDate := getEffectiveDate(scenario.StartDate, i+1)
		projection := buildProjection(
			effectiveDate,
			debts,
			savingsAccounts,
		)
		projections = append(projections, projection)
		if !projection.OutstandingDebt() {
			break
		}

		debts = debt.RefreshDebts(projection.DebtProjections)
		savingsAccounts = savings.RefreshSavingsAccounts(projection.SavingsProjections)
		i = i + 1
	}
	return projections
}

func buildProjection(
	effectiveDate time.Time,
	debts []*debtModels.Debt,
	savingsAccounts []*savingsModels.SavingsAccount,
) *models.Projection {
	debtProjections := debt.ProjectDebtsForMonth(debts, effectiveDate)
	savingsProjections := savings.ProjectSavingsForMonth(savingsAccounts, debtProjections, effectiveDate)

	return &models.Projection{
		EffectiveDate:      effectiveDate,
		DebtProjections:    debtProjections,
		SavingsProjections: savingsProjections,
	}
}

func getEffectiveDate(startDate time.Time, monthCounter int) time.Time {
	return startDate.AddDate(0, monthCounter, 0)
}
