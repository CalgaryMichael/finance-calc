package savings

import (
	models "financeCalc/pkg/savings/models"
	"fmt"
	"reflect"
	"testing"
)

func Test_RefreshSavingsAccounts__NoChange(t *testing.T) {
	payment1 := buildPayment(100.00)
	payment2 := buildPayment(500.00)
	account1 := &models.SavingsAccount{
		Name:           "Jazz 1",
		APY:            0.00,
		InitialCapital: 1400.00,
		Payments:       []*models.SavingsPayment{payment1},
		ProjectedDate:  nil,
	}
	account2 := &models.SavingsAccount{
		Name:           "Jazz 2",
		APY:            0.00,
		InitialCapital: 1500.00,
		Payments:       []*models.SavingsPayment{payment2},
		ProjectedDate:  nil,
	}
	projections := []*models.SavingsProjection{
		&models.SavingsProjection{
			SavingsAccount: account1,
			SavingsTotal:   1400.00,
			PaymentSum:     0.00,
		},
		&models.SavingsProjection{
			SavingsAccount: account2,
			SavingsTotal:   1500.00,
			PaymentSum:     0.00,
		},
	}

	actual := RefreshSavingsAccounts(projections)
	expected := []*models.SavingsAccount{
		account1,
		account2,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Savings Account Refreshes don't match. Actual: %+v; Expected %+v", actual, expected)
		for i, _ := range expected {
			fmt.Printf("[%b] - Actual: %+v - Expected %+v\n", i, actual[i], expected[i])
		}
	}
}

func Test_RefreshSavingsAccounts__TotalChange(t *testing.T) {
	payment1 := buildPayment(100.00)
	payment2 := buildPayment(500.00)
	account1 := &models.SavingsAccount{
		Name:           "Jazz 1",
		APY:            0.00,
		InitialCapital: 1400.00,
		Payments:       []*models.SavingsPayment{payment1},
		ProjectedDate:  nil,
	}
	account2 := &models.SavingsAccount{
		Name:           "Jazz 2",
		APY:            0.00,
		InitialCapital: 1500.00,
		Payments:       []*models.SavingsPayment{payment2},
		ProjectedDate:  nil,
	}
	projections := []*models.SavingsProjection{
		&models.SavingsProjection{
			SavingsAccount: account1,
			SavingsTotal:   1500.00,
			PaymentSum:     100.00,
		},
		&models.SavingsProjection{
			SavingsAccount: account2,
			SavingsTotal:   2000.00,
			PaymentSum:     500.00,
		},
	}

	actual := RefreshSavingsAccounts(projections)
	expected := []*models.SavingsAccount{
		&models.SavingsAccount{
			Name:           "Jazz 1",
			APY:            0.00,
			InitialCapital: 1500.00,
			Payments:       []*models.SavingsPayment{payment1},
			ProjectedDate:  nil,
		},
		&models.SavingsAccount{
			Name:           "Jazz 2",
			APY:            0.00,
			InitialCapital: 2000.00,
			Payments:       []*models.SavingsPayment{payment2},
			ProjectedDate:  nil,
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Savings Account Refreshes don't match. Actual: %+v; Expected %+v", actual, expected)
		for i, _ := range expected {
			fmt.Printf("[%b] - Actual: %+v - Expected %+v\n", i, actual[i], expected[i])
		}
	}
}
