package savings

import (
	"testing"
	"time"
)

func TestIsActive__No_Dates(t *testing.T) {
	payment := SavingsPayment{
		Amount:    100.00,
		StartDate: nil,
		EndDate:   nil,
	}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	actual := payment.IsActive(currentDate)

	if actual != true {
		t.Error("SavingsPayment is inactive; expected to be active")
	}
}

func TestIsActive__Within_Dates(t *testing.T) {
	startDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(1, time.June, 1, 0, 0, 0, 0, time.UTC)
	payment := SavingsPayment{
		Amount:    100.00,
		StartDate: &startDate,
		EndDate:   &endDate,
	}
	currentDate := time.Date(1, time.April, 1, 0, 0, 0, 0, time.UTC)
	actual := payment.IsActive(currentDate)

	if actual != true {
		t.Error("SavingsPayment is inactive; expected to be active")
	}
}

func TestIsActive__Before_Start(t *testing.T) {
	startDate := time.Date(1, time.April, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(1, time.June, 1, 0, 0, 0, 0, time.UTC)
	payment := SavingsPayment{
		Amount:    100.00,
		StartDate: &startDate,
		EndDate:   &endDate,
	}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	actual := payment.IsActive(currentDate)

	if actual != false {
		t.Error("SavingsPayment is active; expected to be inactive")
	}
}

func TestIsActive__After_End(t *testing.T) {
	startDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(1, time.April, 1, 0, 0, 0, 0, time.UTC)
	payment := SavingsPayment{
		Amount:    100.00,
		StartDate: &startDate,
		EndDate:   &endDate,
	}
	currentDate := time.Date(1, time.June, 1, 0, 0, 0, 0, time.UTC)
	actual := payment.IsActive(currentDate)

	if actual != false {
		t.Error("SavingsPayment is active; expected to be inactive")
	}
}
