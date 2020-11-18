package debt

import (
	models "financeCalc/debt/models"
	"reflect"
	"testing"
	"time"
)

func Test_ProjectDebtsForMonth__MultipleDebts(t *testing.T) {
	payments := []*models.DebtPayment{
		&models.DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	debt1 := &models.Debt{
		DebtName:     "Jazz 1",
		DebtTotal:    100.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	debt2 := &models.Debt{
		DebtName:     "Jazz 2",
		DebtTotal:    200.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	debts := []*models.Debt{
		debt1,
		debt2,
	}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)

	actual := ProjectDebtsForMonth(debts, currentDate)
	expected := []*models.DebtProjection{
		&models.DebtProjection{
			Debt:       debt1,
			DebtTotal:  50.00,
			PaymentSum: 50.00,
		},
		&models.DebtProjection{
			Debt:       debt2,
			DebtTotal:  150.00,
			PaymentSum: 50.00,
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("DebtProjections do not match expected. Actual: %+v; Expected: %+v", actual, expected)
	}
}

func Test_BuildProjection(t *testing.T) {
	payments := []*models.DebtPayment{
		&models.DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	debt := &models.Debt{
		DebtName:     "Jazz",
		DebtTotal:    100.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)

	actual := buildProjection(debt, currentDate, 15.00)
	expected := models.DebtProjection{
		Debt:       debt,
		DebtTotal:  35.00,
		PaymentSum: 65.00,
	}

	if *actual != expected {
		t.Errorf("Debt Projection does not match expected. Actual: %+v; Expected: %+v", actual, expected)
	}
}

func Test_GetCarryOverSum__NoDebtProjections(t *testing.T) {
	debts := []*models.DebtProjection{nil, nil}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)

	actual := GetCarryOverSum(debts, currentDate)
	expected := 0.00

	if actual != expected {
		t.Errorf("Carry over from previous debt does not match. Actual: %f; Expected %f", actual, expected)
	}
}

func Test_GetCarryOverSum__WithDebtProjection__NoCarryOver(t *testing.T) {
	payments := []*models.DebtPayment{
		&models.DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	debt := &models.Debt{
		DebtName:     "Jazz",
		DebtTotal:    100.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	debtProjections := []*models.DebtProjection{
		&models.DebtProjection{
			Debt:       debt,
			DebtTotal:  50.00,
			PaymentSum: 50.00,
		},
	}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)

	actual := GetCarryOverSum(debtProjections, currentDate)
	expected := 0.00

	if actual != expected {
		t.Errorf("Carry over from previous debt does not match. Actual: %f; Expected %f", actual, expected)
	}
}

func Test_GetCarryOverSum__WithDebtProjection__WithCarryOver(t *testing.T) {
	payments := []*models.DebtPayment{
		&models.DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	debt1 := &models.Debt{
		DebtName:     "Jazz 1",
		DebtTotal:    0.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	debt2 := &models.Debt{
		DebtName:     "Jazz 2",
		DebtTotal:    50.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	debtProjections := []*models.DebtProjection{
		&models.DebtProjection{
			Debt:       debt1,
			DebtTotal:  0.00,
			PaymentSum: 0.00,
		},
		&models.DebtProjection{
			Debt:       debt2,
			DebtTotal:  0.00,
			PaymentSum: 100.00,
		},
	}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)

	actual := GetCarryOverSum(debtProjections, currentDate)
	expected := 50.00

	if actual != expected {
		t.Errorf("Carry over from previous debt does not match. Actual: %f; Expected %f", actual, expected)
	}
}

func Test_SubtractPaymentSum__AboveThreshold(t *testing.T) {
	actualTotal, actualRemainder := subtractPaymentSum(
		100.00,
		10.00,
		0.12,
	)

	expectedTotal := 91.00
	expectedRemainder := 0.00

	if actualTotal != expectedTotal {
		t.Errorf("Sum of total subtraction does not match. Actual: %.2f; Expected %.2f", actualTotal, expectedTotal)
	}
	if actualRemainder != expectedRemainder {
		t.Errorf("Remainder of subtraction does not match. Actual: %.2f; Expected %.2f", actualRemainder, expectedRemainder)
	}
}

func Test_SubtractPaymentSum__BelowThreshold(t *testing.T) {
	actualTotal, actualRemainder := subtractPaymentSum(
		100.00,
		200.00,
		0.12,
	)

	expectedTotal := 0.00
	expectedRemainder := 99.00

	if actualTotal != expectedTotal {
		t.Errorf("Sum of total subtraction does not match. Actual: %.2f; Expected %.2f", actualTotal, expectedTotal)
	}
	if actualRemainder != expectedRemainder {
		t.Errorf("Remainder of subtraction does not match. Actual: %.2f; Expected %.2f", actualRemainder, expectedRemainder)
	}
}
