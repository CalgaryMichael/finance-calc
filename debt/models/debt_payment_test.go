package debt

import (
	"testing"
	"time"
)

func TestIsActive__No_Dates(t *testing.T) {
	debtPayment := DebtPayment{
		Amount:    100.00,
		CarryOver: true,
		StartDate: nil,
		EndDate:   nil,
	}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	actual := debtPayment.IsActive(currentDate)

	if actual != true {
		t.Error("DebtPayment is inactive; expected to be active")
	}
}

func TestIsActive__Within_Dates(t *testing.T) {
	startDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(1, time.June, 1, 0, 0, 0, 0, time.UTC)
	debtPayment := DebtPayment{
		Amount:    100.00,
		CarryOver: true,
		StartDate: &startDate,
		EndDate:   &endDate,
	}
	currentDate := time.Date(1, time.April, 1, 0, 0, 0, 0, time.UTC)
	actual := debtPayment.IsActive(currentDate)

	if actual != true {
		t.Error("DebtPayment is inactive; expected to be active")
	}
}

func TestIsActive__Before_Start(t *testing.T) {
	startDate := time.Date(1, time.April, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(1, time.June, 1, 0, 0, 0, 0, time.UTC)
	debtPayment := DebtPayment{
		Amount:    100.00,
		CarryOver: true,
		StartDate: &startDate,
		EndDate:   &endDate,
	}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	actual := debtPayment.IsActive(currentDate)

	if actual != false {
		t.Error("DebtPayment is active; expected to be inactive")
	}
}

func TestIsActive__After_End(t *testing.T) {
	startDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(1, time.April, 1, 0, 0, 0, 0, time.UTC)
	debtPayment := DebtPayment{
		Amount:    100.00,
		CarryOver: true,
		StartDate: &startDate,
		EndDate:   &endDate,
	}
	currentDate := time.Date(1, time.June, 1, 0, 0, 0, 0, time.UTC)
	actual := debtPayment.IsActive(currentDate)

	if actual != false {
		t.Error("DebtPayment is active; expected to be inactive")
	}
}
