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

	log.Println("Finished building projections!")
	return projections
}

func buildProjection(
	effectiveDate time.Time,
	debts []*debtModels.Debt,
	savingsAccounts []*savingsModels.SavingsAccount,
) *models.Projection {
	log.Printf("Building projections for month: %v\n", effectiveDate.Format(time.RFC3339))

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
