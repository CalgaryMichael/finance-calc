package savings

import (
	debtModels "financeCalc/debt/models"
	models "financeCalc/savings/models"
	"reflect"
	"testing"
	"time"
)

func Test_ProjectSavingsForMonth__NoCarryOver(t *testing.T) {
	payment1 := buildPayment(100.00)
	payment2 := buildPayment(500.00)
	accounts := []*models.SavingsAccount{
		&models.SavingsAccount{
			Name:           "Jazz 1",
			APY:            0.00,
			InitialCapital: 1400.00,
			Payments:       []*models.SavingsPayment{payment1},
			ProjectedDate:  nil,
		},
		&models.SavingsAccount{
			Name:           "Jazz 2",
			APY:            0.00,
			InitialCapital: 1500.00,
			Payments:       []*models.SavingsPayment{payment2},
			ProjectedDate:  nil,
		},
	}
	debtProjections := []*debtModels.DebtProjection{
		&debtModels.DebtProjection{
			Debt:       nil,
			DebtTotal:  1000.00,
			PaymentSum: 100.00,
		},
	}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)

	actual := ProjectSavingsForMonth(accounts, debtProjections, currentDate)

	expected := []*models.SavingsProjection{
		&models.SavingsProjection{
			SavingsAccount: accounts[0],
			SavingsTotal:   1500.00,
			PaymentSum:     100.00,
		},
		&models.SavingsProjection{
			SavingsAccount: accounts[1],
			SavingsTotal:   2000.00,
			PaymentSum:     500.00,
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("SavingsProjections do not match expected. Actual: %+v; Expected: %+v", actual, expected)
	}
}

func Test_ProjectSavingsForMonth__WithCarryOver(t *testing.T) {
	payment1 := buildPayment(100.00)
	payment2 := buildPayment(500.00)
	accounts := []*models.SavingsAccount{
		&models.SavingsAccount{
			Name:           "Jazz 1",
			APY:            0.00,
			InitialCapital: 1400.00,
			Payments:       []*models.SavingsPayment{payment1},
			ProjectedDate:  nil,
		},
		&models.SavingsAccount{
			Name:           "Jazz 2",
			APY:            0.00,
			InitialCapital: 1500.00,
			Payments:       []*models.SavingsPayment{payment2},
			ProjectedDate:  nil,
		},
	}
	debtProjections := []*debtModels.DebtProjection{
		&debtModels.DebtProjection{
			Debt:       nil,
			DebtTotal:  0.00,
			PaymentSum: 100.00,
		},
	}
	currentDate := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)

	actual := ProjectSavingsForMonth(accounts, debtProjections, currentDate)

	expected := []*models.SavingsProjection{
		&models.SavingsProjection{
			SavingsAccount: accounts[0],
			SavingsTotal:   1600.00,
			PaymentSum:     200.00,
		},
		&models.SavingsProjection{
			SavingsAccount: accounts[1],
			SavingsTotal:   2000.00,
			PaymentSum:     500.00,
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("SavingsProjections do not match expected. Actual: %+v; Expected: %+v", actual, expected)
	}
}

func buildPayment(amount float64) *models.SavingsPayment {
	return &models.SavingsPayment{
		Amount:    amount,
		StartDate: nil,
		EndDate:   nil,
	}
}
