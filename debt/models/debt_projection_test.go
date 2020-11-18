package debt

import (
	"testing"
	"time"
)

func Test_GetActiveSettledPayments__NoSettledDebts(t *testing.T) {
	payments := []*DebtPayment{
		&DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	debt := &Debt{
		DebtName:     "Jazz",
		DebtTotal:    100.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	projection := DebtProjection{
		Debt:       debt,
		DebtTotal:  50.00,
		PaymentSum: 50.00,
	}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)

	actual := projection.GetActiveSettledPayments(currentDate)
	expected := 0.00

	if actual != expected {
		t.Errorf("Active Settled Payments does not match expected. Actual: %.2f; Expected: %.2f", actual, expected)
	}
}

func Test_GetActiveSettledPayments__WithSettled__NoPayment(t *testing.T) {
	payments := []*DebtPayment{
		&DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	debt := &Debt{
		DebtName:     "Jazz",
		DebtTotal:    30.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	projection := DebtProjection{
		Debt:       debt,
		DebtTotal:  0.00,
		PaymentSum: 0.00,
	}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)

	actual := projection.GetActiveSettledPayments(currentDate)
	expected := 20.00

	if actual != expected {
		t.Errorf("Active Settled Payments does not match expected. Actual: %.2f; Expected: %.2f", actual, expected)
	}
}

func Test_GetActiveSettledPayments__WithSettled__WithPayment(t *testing.T) {
	payments := []*DebtPayment{
		&DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	debt := &Debt{
		DebtName:     "Jazz",
		DebtTotal:    30.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	projection := DebtProjection{
		Debt:       debt,
		DebtTotal:  0.00,
		PaymentSum: 100.00,
	}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)

	actual := projection.GetActiveSettledPayments(currentDate)
	expected := 20.00

	if actual != expected {
		t.Errorf("Active Settled Payments does not match expected. Actual: %.2f; Expected: %.2f", actual, expected)
	}
}
