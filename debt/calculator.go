package debt

import (
	models "financeCalc/debt/models"
	"math"
	"time"
)

// Calculate debt reduction cycle for the amount of months between the provided dates
func ProjectDebtsForMonth(debts []*models.Debt, endDate time.Time) []*models.DebtProjection {
	projections := make([]*models.DebtProjection, len(debts))
	for i, debt := range debts {
		carryOverSum := GetCarryOverSum(projections)
		projections[i] = buildProjection(debt, endDate, carryOverSum)
	}
	return projections
}

func buildProjection(debt *models.Debt, currentDate time.Time, carryOverSum float64) *models.DebtProjection {
	paymentSum := debt.SumActivePayments(currentDate) + carryOverSum
	debtTotal := subtractPaymentSum(debt.DebtTotal, paymentSum, debt.InterestRate)
	return &models.DebtProjection{
		Debt:       debt,
		DebtTotal:  debtTotal,
		PaymentSum: paymentSum,
	}
}

func GetCarryOverSum(projections []*models.DebtProjection) float64 {
	projectionCount := len(projections)
	if projectionCount == 0 {
		return 0.00
	}
	lastProjection := getLastProjection(projections)
	if lastProjection == nil {
		return 0.00
	}
	return math.Abs(math.Min(lastProjection.DebtTotal-lastProjection.PaymentSum, 0.00))
}

func getLastProjection(projections []*models.DebtProjection) *models.DebtProjection {
	var projection *models.DebtProjection
	for i := len(projections) - 1; i >= 0; i-- {
		if projections[i] != nil {
			projection = projections[i]
			break
		}
	}
	return projection
}

func subtractPaymentSum(debtTotal float64, paymentSum float64, interestRate float64) float64 {
	debtTotal = debtTotal * (1 + interestRate/12)
	return math.Max(debtTotal-paymentSum, 0.00)
}
