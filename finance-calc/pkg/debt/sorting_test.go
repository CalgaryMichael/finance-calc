package debt

import (
	models "financeCalc/pkg/debt/models"
	"reflect"
	"testing"
)

func Test_SortDebts__SettledOnTop(t *testing.T) {
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
		DebtTotal:    0.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	debts := []*models.Debt{
		debt1,
		debt2,
	}

	actual := SortDebts(debts, "DebtTotal", false)

	expected := []*models.Debt{
		debt2,
		debt1,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Debt Refreshes don't match. Actual: %+v; Expected %+v", actual, expected)
	}
}

func Test_SortDebts__RespectsSortKey__DebtName(t *testing.T) {
	payments := []*models.DebtPayment{
		&models.DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	debt1 := &models.Debt{
		DebtName:     "DEF",
		DebtTotal:    100.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	debt2 := &models.Debt{
		DebtName:     "ABC",
		DebtTotal:    100.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	debts := []*models.Debt{
		debt1,
		debt2,
	}

	actual := SortDebts(debts, "DebtName", false)

	expected := []*models.Debt{
		debt2,
		debt1,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Debt Refreshes don't match. Actual: %+v; Expected %+v", actual, expected)
	}
}

func Test_SortDebts__RespectsSortKey__InterestRate(t *testing.T) {
	payments := []*models.DebtPayment{
		&models.DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	debt1 := &models.Debt{
		DebtName:     "Jazz",
		DebtTotal:    100.00,
		Payments:     payments,
		InterestRate: 0.20,
	}
	debt2 := &models.Debt{
		DebtName:     "Jazz",
		DebtTotal:    100.00,
		Payments:     payments,
		InterestRate: 0.10,
	}
	debts := []*models.Debt{
		debt1,
		debt2,
	}

	actual := SortDebts(debts, "InterestRate", false)

	expected := []*models.Debt{
		debt2,
		debt1,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Debt Refreshes don't match. Actual: %+v; Expected %+v", actual, expected)
	}
}

func Test_SortDebts__RespectsSortKey__Payments(t *testing.T) {
	payments1 := []*models.DebtPayment{
		&models.DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	payments2 := []*models.DebtPayment{
		&models.DebtPayment{
			Amount:    75.00,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	debt1 := &models.Debt{
		DebtName:     "ABC",
		DebtTotal:    100.00,
		Payments:     payments1,
		InterestRate: 0.00,
	}
	debt2 := &models.Debt{
		DebtName:     "ABC",
		DebtTotal:    100.00,
		Payments:     payments2,
		InterestRate: 0.00,
	}
	debts := []*models.Debt{
		debt1,
		debt2,
	}

	actual := SortDebts(debts, "Payments", false)

	expected := []*models.Debt{
		debt2,
		debt1,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Debt Refreshes don't match. Actual: %+v; Expected %+v", actual, expected)
	}
}

func Test_SortDebts__RespectsReversed(t *testing.T) {
	payments := []*models.DebtPayment{
		&models.DebtPayment{
			Amount:    50.00,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	debt1 := &models.Debt{
		DebtName:     "DEF",
		DebtTotal:    100.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	debt2 := &models.Debt{
		DebtName:     "ABC",
		DebtTotal:    100.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	debts := []*models.Debt{
		debt2,
		debt1,
	}

	actual := SortDebts(debts, "DebtName", true)

	expected := []*models.Debt{
		debt1,
		debt2,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Debt Refreshes don't match. Actual: %+v; Expected %+v", actual, expected)
	}
}

func Test_SortSettledDebts__SettledOnTop(t *testing.T) {
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
		DebtTotal:    0.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	debt3 := &models.Debt{
		DebtName:     "Jazz 3",
		DebtTotal:    10.00,
		Payments:     payments,
		InterestRate: 0.00,
	}
	debts := []*models.Debt{
		debt1,
		debt2,
		debt3,
	}

	actual := SortSettledDebts(debts)

	expected := []*models.Debt{
		debt2,
		debt1,
		debt3,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Debt Refreshes don't match. Actual: %+v; Expected %+v", actual, expected)
	}
}
