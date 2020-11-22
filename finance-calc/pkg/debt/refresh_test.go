package debt

import (
	models "financeCalc/pkg/debt/models"
	"reflect"
	"testing"
)

func Test_RefreshDebts__NoChange(t *testing.T) {
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

	projections := []*models.DebtProjection{
		&models.DebtProjection{
			Debt:       debt1,
			DebtTotal:  100.00,
			PaymentSum: 0.00,
		},
		&models.DebtProjection{
			Debt:       debt2,
			DebtTotal:  200.00,
			PaymentSum: 0.00,
		},
	}

	actual := RefreshDebts(projections)
	expected := []*models.Debt{
		debt1,
		debt2,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Debt Refreshes don't match. Actual: %+v; Expected %+v", actual, expected)
	}
}

func Test_RefreshDebts__DebtChange(t *testing.T) {
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

	projections := []*models.DebtProjection{
		&models.DebtProjection{
			Debt:       debt1,
			DebtTotal:  75.00,
			PaymentSum: 25.00,
		},
		&models.DebtProjection{
			Debt:       debt2,
			DebtTotal:  150.00,
			PaymentSum: 50.00,
		},
	}

	actual := RefreshDebts(projections)

	expected := []*models.Debt{
		&models.Debt{
			DebtName:     "Jazz 1",
			DebtTotal:    75.00,
			Payments:     payments,
			InterestRate: 0.00,
		},
		&models.Debt{
			DebtName:     "Jazz 2",
			DebtTotal:    150.00,
			Payments:     payments,
			InterestRate: 0.00,
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Debt Refreshes don't match. Actual: %+v; Expected %+v", actual, expected)
	}
}

func Test_RefreshDebts__SettledOnTop(t *testing.T) {
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

	projections := []*models.DebtProjection{
		&models.DebtProjection{
			Debt:       debt1,
			DebtTotal:  75.00,
			PaymentSum: 25.00,
		},
		&models.DebtProjection{
			Debt:       debt2,
			DebtTotal:  0.00,
			PaymentSum: 50.00,
		},
	}

	actual := RefreshDebts(projections)

	expected := []*models.Debt{
		&models.Debt{
			DebtName:     "Jazz 2",
			DebtTotal:    0.00,
			Payments:     payments,
			InterestRate: 0.00,
		},
		&models.Debt{
			DebtName:     "Jazz 1",
			DebtTotal:    75.00,
			Payments:     payments,
			InterestRate: 0.00,
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Debt Refreshes don't match. Actual: %+v; Expected %+v", actual, expected)
	}
}
