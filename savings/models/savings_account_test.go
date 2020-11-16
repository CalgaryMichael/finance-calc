package savings

import (
	"testing"
	"time"
)

func TestSumActivePayments__AllActive__NoDates(t *testing.T) {
	payments := []*SavingsPayment{
		&SavingsPayment{
			Amount:    50.00,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	account := SavingsAccount{
		Name:           "Jazz",
		APY:            0.00,
		InitialCapital: 100.00,
		Payments:       payments,
	}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)

	actual := account.SumActivePayments(currentDate)
	expected := 50.00

	if actual != expected {
		t.Errorf("Sum of active payments is incorrect. Actual: %.2f; Expected %.2f", actual, expected)
	}
}

func TestSumActivePayments__AllActive__WithDates(t *testing.T) {
	startDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(1, time.June, 1, 0, 0, 0, 0, time.UTC)
	payments := []*SavingsPayment{
		&SavingsPayment{
			Amount:    50.00,
			StartDate: &startDate,
			EndDate:   &endDate,
		},
	}
	account := SavingsAccount{
		Name:           "Jazz",
		APY:            0.00,
		InitialCapital: 100.00,
		Payments:       payments,
	}
	currentDate := time.Date(1, time.April, 1, 0, 0, 0, 0, time.UTC)

	actual := account.SumActivePayments(currentDate)
	expected := 50.00

	if actual != expected {
		t.Errorf("Sum of active payments is incorrect. Actual: %.2f; Expected %.2f", actual, expected)
	}
}

func TestSumActivePayments__NoneActive(t *testing.T) {
	startDate := time.Date(1, time.April, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(1, time.June, 1, 0, 0, 0, 0, time.UTC)
	payments := []*SavingsPayment{
		&SavingsPayment{
			Amount:    50.00,
			StartDate: &startDate,
			EndDate:   &endDate,
		},
	}
	account := SavingsAccount{
		Name:           "Jazz",
		APY:            0.00,
		InitialCapital: 100.00,
		Payments:       payments,
	}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)

	actual := account.SumActivePayments(currentDate)
	expected := 0.00

	if actual != expected {
		t.Errorf("Sum of active payments is incorrect. Actual: %.2f; Expected %.2f", actual, expected)
	}
}

func TestSumActivePayments__Multiple(t *testing.T) {
	startDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(1, time.June, 1, 0, 0, 0, 0, time.UTC)
	payments := []*SavingsPayment{
		&SavingsPayment{
			Amount:    50.00,
			StartDate: &startDate,
			EndDate:   &endDate,
		},
		&SavingsPayment{
			Amount:    25.00,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	account := SavingsAccount{
		Name:           "Jazz",
		APY:            0.00,
		InitialCapital: 100.00,
		Payments:       payments,
	}
	currentDate := time.Date(1, time.April, 1, 0, 0, 0, 0, time.UTC)

	actual := account.SumActivePayments(currentDate)
	expected := 75.00

	if actual != expected {
		t.Errorf("Sum of active payments is incorrect. Actual: %.2f; Expected %.2f", actual, expected)
	}
}

func TestSumActivePayments__Multiple__Mixed(t *testing.T) {
	startDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(1, time.June, 1, 0, 0, 0, 0, time.UTC)
	payments := []*SavingsPayment{
		&SavingsPayment{
			Amount:    50.00,
			StartDate: &startDate,
			EndDate:   &endDate,
		},
		&SavingsPayment{
			Amount:    25.00,
			StartDate: &endDate,
			EndDate:   nil,
		},
	}
	account := SavingsAccount{
		Name:           "Jazz",
		APY:            0.00,
		InitialCapital: 100.00,
		Payments:       payments,
	}
	currentDate := time.Date(1, time.April, 1, 0, 0, 0, 0, time.UTC)

	actual := account.SumActivePayments(currentDate)
	expected := 50.00

	if actual != expected {
		t.Errorf("Sum of active payments is incorrect. Actual: %.2f; Expected %.2f", actual, expected)
	}
}
