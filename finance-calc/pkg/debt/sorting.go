package debt

import (
	models "financeCalc/pkg/debt/models"
	"sort"
)

func SortDebts(debts []*models.Debt, key string, reverse bool) []*models.Debt {
	sort.Slice(
		debts,
		func(i, j int) bool {
			return cmpSettledDebts(debts[i], debts[j]) || cmpDebts(debts[i], debts[j], key, reverse)
		},
	)
	return debts
}

func SortSettledDebts(debts []*models.Debt) []*models.Debt {
	sort.Slice(
		debts,
		func(i, j int) bool {
			return cmpSettledDebts(debts[i], debts[j])
		},
	)
	return debts
}

func cmpSettledDebts(first *models.Debt, second *models.Debt) bool {
	return first.IsSettled() && !second.IsSettled()
}

func cmpDebts(first *models.Debt, second *models.Debt, key string, reverse bool) bool {
	var comparison bool
	switch key {
	case "DebtName":
		comparison = first.DebtName < second.DebtName
	case "InterestRate":
		comparison = first.InterestRate < second.InterestRate
	case "Payments":
		comparison = first.SumPayments() > second.SumPayments()
	default:
		comparison = first.DebtTotal < second.DebtTotal
	}

	return (comparison && !reverse) || (!comparison && reverse)
}
