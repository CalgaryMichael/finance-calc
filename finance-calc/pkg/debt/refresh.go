package debt

import models "financeCalc/pkg/debt/models"

func RefreshDebts(projections []*models.DebtProjection) []*models.Debt {
	debts := make([]*models.Debt, len(projections))
	for i, projection := range projections {
		debts[i] = copyDebt(projection)
	}
	return SortSettledDebts(debts)
}

func copyDebt(projection *models.DebtProjection) *models.Debt {
	return &models.Debt{
		DebtName:     projection.Debt.DebtName,
		DebtTotal:    projection.DebtTotal,
		Payments:     projection.Debt.Payments,
		InterestRate: projection.Debt.InterestRate,
	}
}
