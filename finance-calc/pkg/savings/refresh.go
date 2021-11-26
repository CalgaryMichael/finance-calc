package savings

import models "financeCalc/pkg/savings/models"

func RefreshSavingsAccounts(projections []*models.SavingsProjection) []*models.SavingsAccount {
	accounts := make([]*models.SavingsAccount, len(projections))
	for i, projection := range projections {
		accounts[i] = copySavingsAccount(projection)
	}
	return accounts
}

func copySavingsAccount(projection *models.SavingsProjection) *models.SavingsAccount {
	return &models.SavingsAccount{
		Id:             projection.SavingsAccount.Id,
		Name:           projection.SavingsAccount.Name,
		APY:            projection.SavingsAccount.APY,
		InitialCapital: projection.SavingsTotal,
		Payments:       projection.SavingsAccount.Payments,
		ProjectedDate:  projection.SavingsAccount.ProjectedDate,
	}
}
