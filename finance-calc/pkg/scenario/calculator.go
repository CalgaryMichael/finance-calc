package scenario

import (
	debt "financeCalc/pkg/debt"
	debtModels "financeCalc/pkg/debt/models"
	savings "financeCalc/pkg/savings"
	savingsModels "financeCalc/pkg/savings/models"
	models "financeCalc/pkg/scenario/models"
	"log"
	"time"
)

func BuildProjections(scenario models.Scenario) []*models.Projection {
	log.Println("Building projections...")
	projections := make([]*models.Projection, 0, 240)

	debts := debt.SortDebts(scenario.Debts, scenario.SortKey, scenario.ReverseSort)
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
		if !projection.OutstandingDebt() && !savings.OutstandingSavingsProjections(savingsAccounts, effectiveDate) {
			break
		}

		debts = debt.RefreshDebts(projection.DebtProjections)
		savingsAccounts = savings.RefreshSavingsAccounts(projection.SavingsProjections)
		i = i + 1
	}

	log.Println("Finished building projections!")
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
