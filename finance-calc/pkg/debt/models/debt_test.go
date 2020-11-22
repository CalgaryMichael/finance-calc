package debt

import (
	"testing"
	"time"
)

func TestIsSettled(t *testing.T) {
	payments := []*DebtPayment{
		&DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	debt := Debt{
		DebtName:     "Jazz",
		DebtTotal:    0.00,
		Payments:     payments,
		InterestRate: 0.00,
	}

	actual := debt.IsSettled()
	expected := true

	if actual != expected {
		t.Errorf("'IsSettled' returned unexpected value. Actual: %t; Expected %t", actual, expected)
	}
}

func TestIsSettled__Unsettled(t *testing.T) {
	payments := []*DebtPayment{
		&DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	debt := Debt{
		DebtName:     "Jazz",
		DebtTotal:    100.00,
		Payments:     payments,
		InterestRate: 0.00,
	}

	actual := debt.IsSettled()
	expected := false

	if actual != expected {
		t.Errorf("'IsSettled' returned unexpected value. Actual: %t; Expected %t", actual, expected)
	}
}

func TestSumPayments(t *testing.T) {
	payments := []*DebtPayment{
		&DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
		&DebtPayment{
			Amount:    25.00,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	debt := Debt{
		DebtName:     "Jazz",
		DebtTotal:    100.00,
		Payments:     payments,
		InterestRate: 0.00,
	}

	actual := debt.SumPayments()
	expected := 75.00

	if actual != expected {
		t.Errorf("Sum of payments is incorrect. Actual: %.2f; Expected %.2f", actual, expected)
	}
}

func TestSumActivePayments__AllActive__NoDates(t *testing.T) {
	payments := []*DebtPayment{
		&DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	debt := Debt{
		DebtName:     "Jazz",
		DebtTotal:    100.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)

	actual := debt.SumActivePayments(currentDate)
	expected := 50.00

	if actual != expected {
		t.Errorf("Sum of active payments is incorrect. Actual: %.2f; Expected %.2f", actual, expected)
	}
}

func TestSumActivePayments__AllActive__NoCarryOver(t *testing.T) {
	payments := []*DebtPayment{
		&DebtPayment{
			Amount:    50.00,
			CarryOver: false,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	debt := Debt{
		DebtName:     "Jazz",
		DebtTotal:    0.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)

	actual := debt.SumActivePayments(currentDate)
	expected := 0.00

	if actual != expected {
		t.Errorf("Sum of active payments is incorrect. Actual: %.2f; Expected %.2f", actual, expected)
	}
}

func TestSumActivePayments__AllActive__WithDates(t *testing.T) {
	startDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(1, time.June, 1, 0, 0, 0, 0, time.UTC)
	payments := []*DebtPayment{
		&DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: &startDate,
			EndDate:   &endDate,
		},
	}
	debt := Debt{
		DebtName:     "Jazz",
		DebtTotal:    100.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	currentDate := time.Date(1, time.April, 1, 0, 0, 0, 0, time.UTC)

	actual := debt.SumActivePayments(currentDate)
	expected := 50.00

	if actual != expected {
		t.Errorf("Sum of active payments is incorrect. Actual: %.2f; Expected %.2f", actual, expected)
	}
}

func TestSumActivePayments__NoneActive(t *testing.T) {
	startDate := time.Date(1, time.April, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(1, time.June, 1, 0, 0, 0, 0, time.UTC)
	payments := []*DebtPayment{
		&DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: &startDate,
			EndDate:   &endDate,
		},
	}
	debt := Debt{
		DebtName:     "Jazz",
		DebtTotal:    100.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)

	actual := debt.SumActivePayments(currentDate)
	expected := 0.00

	if actual != expected {
		t.Errorf("Sum of active payments is incorrect. Actual: %.2f; Expected %.2f", actual, expected)
	}
}

func TestSumActivePayments__Multiple(t *testing.T) {
	startDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(1, time.June, 1, 0, 0, 0, 0, time.UTC)
	payments := []*DebtPayment{
		&DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: &startDate,
			EndDate:   &endDate,
		},
		&DebtPayment{
			Amount:    25.00,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	debt := Debt{
		DebtName:     "Jazz",
		DebtTotal:    100.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	currentDate := time.Date(1, time.April, 1, 0, 0, 0, 0, time.UTC)

	actual := debt.SumActivePayments(currentDate)
	expected := 75.00

	if actual != expected {
		t.Errorf("Sum of active payments is incorrect. Actual: %.2f; Expected %.2f", actual, expected)
	}
}

func TestSumActivePayments__Multiple__Mixed(t *testing.T) {
	startDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(1, time.June, 1, 0, 0, 0, 0, time.UTC)
	payments := []*DebtPayment{
		&DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: &startDate,
			EndDate:   &endDate,
		},
		&DebtPayment{
			Amount:    25.00,
			CarryOver: true,
			StartDate: &endDate,
			EndDate:   nil,
		},
	}
	debt := Debt{
		DebtName:     "Jazz",
		DebtTotal:    100.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	currentDate := time.Date(1, time.April, 1, 0, 0, 0, 0, time.UTC)

	actual := debt.SumActivePayments(currentDate)
	expected := 50.00

	if actual != expected {
		t.Errorf("Sum of active payments is incorrect. Actual: %.2f; Expected %.2f", actual, expected)
	}
}
