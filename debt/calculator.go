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
		carryOverSum := GetCarryOverSum(projections, endDate)
		projections[i] = buildProjection(debt, endDate, carryOverSum)
	}
	return projections
}

func buildProjection(debt *models.Debt, currentDate time.Time, carryOverSum float64) *models.DebtProjection {
	paymentSum := debt.SumActivePayments(currentDate) + carryOverSum
	debtTotal, remainder := subtractPaymentSum(debt.DebtTotal, paymentSum, debt.InterestRate)
	return &models.DebtProjection{
		Debt:       debt,
		DebtTotal:  debtTotal,
		PaymentSum: paymentSum - remainder,
	}
}

func GetCarryOverSum(projections []*models.DebtProjection, currentDate time.Time) float64 {
	projectionCount := len(projections)
	if projectionCount == 0 {
		return 0.00
	}

	sum := 0.00
	for i := projectionCount - 1; i >= 0; i-- {
		if projections[i] == nil {
			continue
		}
		if !projections[i].IsSettled() {
			sum = 0.00
			break
		}
		sum += projections[i].GetActiveSettledPayments(currentDate)
	}
	return sum
}

func subtractPaymentSum(debtTotal float64, paymentSum float64, interestRate float64) (float64, float64) {
	debtTotal = debtTotal * (1 + interestRate/12)
	return math.Max(debtTotal-paymentSum, 0.00), math.Max(paymentSum-debtTotal, 0.00)
}
