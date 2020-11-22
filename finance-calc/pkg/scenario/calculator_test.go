package scenario

import (
	debtModels "financeCalc/pkg/debt/models"
	savingsModels "financeCalc/pkg/savings/models"
	models "financeCalc/pkg/scenario/models"
	"reflect"
	"testing"
	"time"
)

func Test_BuildProjections__DebtSettled(t *testing.T) {
	debt := buildDebt(100.00, 25.00)
	savingsAccount := buildSavingsAccount(1000.00, 100.00)
	scenario := models.Scenario{
		StartDate:       time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		Debts:           []*debtModels.Debt{debt},
		SavingsAccounts: []*savingsModels.SavingsAccount{savingsAccount},
	}
	actual := BuildProjections(scenario, "DebtTotal", false)

	expected := []*models.Projection{
		&models.Projection{
			EffectiveDate: time.Date(1, time.February, 1, 0, 0, 0, 0, time.UTC),
			DebtProjections: []*debtModels.DebtProjection{
				&debtModels.DebtProjection{
					Debt:         debt,
					DebtTotal:    75.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
			},
			SavingsProjections: []*savingsModels.SavingsProjection{
				&savingsModels.SavingsProjection{
					SavingsAccount: savingsAccount,
					SavingsTotal:   1100.00,
					PaymentSum:     100.00,
				},
			},
		},
		&models.Projection{
			EffectiveDate: time.Date(1, time.March, 1, 0, 0, 0, 0, time.UTC),
			DebtProjections: []*debtModels.DebtProjection{
				&debtModels.DebtProjection{
					Debt:         buildDebt(75.00, 25.00),
					DebtTotal:    50.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
			},
			SavingsProjections: []*savingsModels.SavingsProjection{
				&savingsModels.SavingsProjection{
					SavingsAccount: buildSavingsAccount(1100.00, 100.00),
					SavingsTotal:   1200.00,
					PaymentSum:     100.00,
				},
			},
		},
		&models.Projection{
			EffectiveDate: time.Date(1, time.April, 1, 0, 0, 0, 0, time.UTC),
			DebtProjections: []*debtModels.DebtProjection{
				&debtModels.DebtProjection{
					Debt:         buildDebt(50.00, 25.00),
					DebtTotal:    25.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
			},
			SavingsProjections: []*savingsModels.SavingsProjection{
				&savingsModels.SavingsProjection{
					SavingsAccount: buildSavingsAccount(1200.00, 100.00),
					SavingsTotal:   1300.00,
					PaymentSum:     100.00,
				},
			},
		},
		&models.Projection{
			EffectiveDate: time.Date(1, time.May, 1, 0, 0, 0, 0, time.UTC),
			DebtProjections: []*debtModels.DebtProjection{
				&debtModels.DebtProjection{
					Debt:         buildDebt(25.00, 25.00),
					DebtTotal:    0.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
			},
			SavingsProjections: []*savingsModels.SavingsProjection{
				&savingsModels.SavingsProjection{
					SavingsAccount: buildSavingsAccount(1300.00, 100.00),
					SavingsTotal:   1400.00,
					PaymentSum:     100.00,
				},
			},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Savings Account Refreshes don't match. Actual: %+v; Expected %+v", actual, expected)
	}
}

func Test_BuildProjections__DebtSettled__WaterfallCarryOver(t *testing.T) {
	debt1 := buildDebt(100.00, 25.00)
	debt2 := buildDebt(175.00, 25.00)
	debt3 := buildDebt(250.00, 25.00)
	savingsAccount := buildSavingsAccount(1000.00, 100.00)
	scenario := models.Scenario{
		StartDate:       time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		Debts:           []*debtModels.Debt{debt1, debt2, debt3},
		SavingsAccounts: []*savingsModels.SavingsAccount{savingsAccount},
	}
	actual := BuildProjections(scenario, "DebtTotal", false)

	expected := []*models.Projection{
		&models.Projection{
			EffectiveDate: time.Date(1, time.February, 1, 0, 0, 0, 0, time.UTC),
			DebtProjections: []*debtModels.DebtProjection{
				&debtModels.DebtProjection{
					Debt:         debt1,
					DebtTotal:    75.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
				&debtModels.DebtProjection{
					Debt:         debt2,
					DebtTotal:    150.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
				&debtModels.DebtProjection{
					Debt:         debt3,
					DebtTotal:    225.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
			},
			SavingsProjections: []*savingsModels.SavingsProjection{
				&savingsModels.SavingsProjection{
					SavingsAccount: savingsAccount,
					SavingsTotal:   1100.00,
					PaymentSum:     100.00,
				},
			},
		},
		&models.Projection{
			EffectiveDate: time.Date(1, time.March, 1, 0, 0, 0, 0, time.UTC),
			DebtProjections: []*debtModels.DebtProjection{
				&debtModels.DebtProjection{
					Debt:         buildDebt(75.00, 25.00),
					DebtTotal:    50.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
				&debtModels.DebtProjection{
					Debt:         buildDebt(150.00, 25.00),
					DebtTotal:    125.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
				&debtModels.DebtProjection{
					Debt:         buildDebt(225.00, 25.00),
					DebtTotal:    200.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
			},
			SavingsProjections: []*savingsModels.SavingsProjection{
				&savingsModels.SavingsProjection{
					SavingsAccount: buildSavingsAccount(1100.00, 100.00),
					SavingsTotal:   1200.00,
					PaymentSum:     100.00,
				},
			},
		},
		&models.Projection{
			EffectiveDate: time.Date(1, time.April, 1, 0, 0, 0, 0, time.UTC),
			DebtProjections: []*debtModels.DebtProjection{
				&debtModels.DebtProjection{
					Debt:         buildDebt(50.00, 25.00),
					DebtTotal:    25.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
				&debtModels.DebtProjection{
					Debt:         buildDebt(125.00, 25.00),
					DebtTotal:    100.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
				&debtModels.DebtProjection{
					Debt:         buildDebt(200.00, 25.00),
					DebtTotal:    175.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
			},
			SavingsProjections: []*savingsModels.SavingsProjection{
				&savingsModels.SavingsProjection{
					SavingsAccount: buildSavingsAccount(1200.00, 100.00),
					SavingsTotal:   1300.00,
					PaymentSum:     100.00,
				},
			},
		},
		&models.Projection{
			EffectiveDate: time.Date(1, time.May, 1, 0, 0, 0, 0, time.UTC),
			DebtProjections: []*debtModels.DebtProjection{
				&debtModels.DebtProjection{
					Debt:         buildDebt(25.00, 25.00),
					DebtTotal:    0.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
				&debtModels.DebtProjection{
					Debt:         buildDebt(100.00, 25.00),
					DebtTotal:    75.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
				&debtModels.DebtProjection{
					Debt:         buildDebt(175.00, 25.00),
					DebtTotal:    150.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
			},
			SavingsProjections: []*savingsModels.SavingsProjection{
				&savingsModels.SavingsProjection{
					SavingsAccount: buildSavingsAccount(1300.00, 100.00),
					SavingsTotal:   1400.00,
					PaymentSum:     100.00,
				},
			},
		},
		&models.Projection{
			EffectiveDate: time.Date(1, time.June, 1, 0, 0, 0, 0, time.UTC),
			DebtProjections: []*debtModels.DebtProjection{
				&debtModels.DebtProjection{
					Debt:         buildDebt(0.00, 25.00),
					DebtTotal:    0.00,
					PaymentSum:   0.00,
					UnappliedSum: 25.00,
				},
				&debtModels.DebtProjection{
					Debt:         buildDebt(75.00, 25.00),
					DebtTotal:    25.00,
					PaymentSum:   50.00,
					UnappliedSum: 0.00,
				},
				&debtModels.DebtProjection{
					Debt:         buildDebt(150.00, 25.00),
					DebtTotal:    125.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
			},
			SavingsProjections: []*savingsModels.SavingsProjection{
				&savingsModels.SavingsProjection{
					SavingsAccount: buildSavingsAccount(1400.00, 100.00),
					SavingsTotal:   1500.00,
					PaymentSum:     100.00,
				},
			},
		},
		&models.Projection{
			EffectiveDate: time.Date(1, time.July, 1, 0, 0, 0, 0, time.UTC),
			DebtProjections: []*debtModels.DebtProjection{
				&debtModels.DebtProjection{
					Debt:         buildDebt(0.00, 25.00),
					DebtTotal:    0.00,
					PaymentSum:   0.00,
					UnappliedSum: 25.00,
				},
				&debtModels.DebtProjection{
					Debt:         buildDebt(25.00, 25.00),
					DebtTotal:    0.00,
					PaymentSum:   25.00,
					UnappliedSum: 25.00,
				},
				&debtModels.DebtProjection{
					Debt:         buildDebt(125.00, 25.00),
					DebtTotal:    75.00,
					PaymentSum:   50.00,
					UnappliedSum: 0.00,
				},
			},
			SavingsProjections: []*savingsModels.SavingsProjection{
				&savingsModels.SavingsProjection{
					SavingsAccount: buildSavingsAccount(1500.00, 100.00),
					SavingsTotal:   1600.00,
					PaymentSum:     100.00,
				},
			},
		},
		&models.Projection{
			EffectiveDate: time.Date(1, time.August, 1, 0, 0, 0, 0, time.UTC),
			DebtProjections: []*debtModels.DebtProjection{
				&debtModels.DebtProjection{
					Debt:         buildDebt(0.00, 25.00),
					DebtTotal:    0.00,
					PaymentSum:   0.00,
					UnappliedSum: 25.00,
				},
				&debtModels.DebtProjection{
					Debt:         buildDebt(0.00, 25.00),
					DebtTotal:    0.00,
					PaymentSum:   0.00,
					UnappliedSum: 50.00,
				},
				&debtModels.DebtProjection{
					Debt:         buildDebt(75.00, 25.00),
					DebtTotal:    0.00,
					PaymentSum:   75.00,
					UnappliedSum: 0.00,
				},
			},
			SavingsProjections: []*savingsModels.SavingsProjection{
				&savingsModels.SavingsProjection{
					SavingsAccount: buildSavingsAccount(1600.00, 100.00),
					SavingsTotal:   1700.00,
					PaymentSum:     100.00,
				},
			},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Savings Account Refreshes don't match. Actual: %+v; Expected %+v", actual, expected)
	}
}

func Test_BuildProjections__DebtSettled__SavingsCarryOver(t *testing.T) {
	debt1 := buildDebt(100.00, 25.00)
	debt2 := buildDebt(175.00, 25.00)
	savingsAccount := buildSavingsAccount(1000.00, 100.00)
	scenario := models.Scenario{
		StartDate:       time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		Debts:           []*debtModels.Debt{debt1, debt2},
		SavingsAccounts: []*savingsModels.SavingsAccount{savingsAccount},
	}
	actual := BuildProjections(scenario, "DebtTotal", false)

	expected := []*models.Projection{
		&models.Projection{
			EffectiveDate: time.Date(1, time.February, 1, 0, 0, 0, 0, time.UTC),
			DebtProjections: []*debtModels.DebtProjection{
				&debtModels.DebtProjection{
					Debt:         debt1,
					DebtTotal:    75.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
				&debtModels.DebtProjection{
					Debt:         debt2,
					DebtTotal:    150.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
			},
			SavingsProjections: []*savingsModels.SavingsProjection{
				&savingsModels.SavingsProjection{
					SavingsAccount: savingsAccount,
					SavingsTotal:   1100.00,
					PaymentSum:     100.00,
				},
			},
		},
		&models.Projection{
			EffectiveDate: time.Date(1, time.March, 1, 0, 0, 0, 0, time.UTC),
			DebtProjections: []*debtModels.DebtProjection{
				&debtModels.DebtProjection{
					Debt:         buildDebt(75.00, 25.00),
					DebtTotal:    50.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
				&debtModels.DebtProjection{
					Debt:         buildDebt(150.00, 25.00),
					DebtTotal:    125.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
			},
			SavingsProjections: []*savingsModels.SavingsProjection{
				&savingsModels.SavingsProjection{
					SavingsAccount: buildSavingsAccount(1100.00, 100.00),
					SavingsTotal:   1200.00,
					PaymentSum:     100.00,
				},
			},
		},
		&models.Projection{
			EffectiveDate: time.Date(1, time.April, 1, 0, 0, 0, 0, time.UTC),
			DebtProjections: []*debtModels.DebtProjection{
				&debtModels.DebtProjection{
					Debt:         buildDebt(50.00, 25.00),
					DebtTotal:    25.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
				&debtModels.DebtProjection{
					Debt:         buildDebt(125.00, 25.00),
					DebtTotal:    100.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
			},
			SavingsProjections: []*savingsModels.SavingsProjection{
				&savingsModels.SavingsProjection{
					SavingsAccount: buildSavingsAccount(1200.00, 100.00),
					SavingsTotal:   1300.00,
					PaymentSum:     100.00,
				},
			},
		},
		&models.Projection{
			EffectiveDate: time.Date(1, time.May, 1, 0, 0, 0, 0, time.UTC),
			DebtProjections: []*debtModels.DebtProjection{
				&debtModels.DebtProjection{
					Debt:         buildDebt(25.00, 25.00),
					DebtTotal:    0.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
				&debtModels.DebtProjection{
					Debt:         buildDebt(100.00, 25.00),
					DebtTotal:    75.00,
					PaymentSum:   25.00,
					UnappliedSum: 0.00,
				},
			},
			SavingsProjections: []*savingsModels.SavingsProjection{
				&savingsModels.SavingsProjection{
					SavingsAccount: buildSavingsAccount(1300.00, 100.00),
					SavingsTotal:   1400.00,
					PaymentSum:     100.00,
				},
			},
		},
		&models.Projection{
			EffectiveDate: time.Date(1, time.June, 1, 0, 0, 0, 0, time.UTC),
			DebtProjections: []*debtModels.DebtProjection{
				&debtModels.DebtProjection{
					Debt:         buildDebt(0.00, 25.00),
					DebtTotal:    0.00,
					PaymentSum:   0.00,
					UnappliedSum: 25.00,
				},
				&debtModels.DebtProjection{
					Debt:         buildDebt(75.00, 25.00),
					DebtTotal:    25.00,
					PaymentSum:   50.00,
					UnappliedSum: 0.00,
				},
			},
			SavingsProjections: []*savingsModels.SavingsProjection{
				&savingsModels.SavingsProjection{
					SavingsAccount: buildSavingsAccount(1400.00, 100.00),
					SavingsTotal:   1500.00,
					PaymentSum:     100.00,
				},
			},
		},
		&models.Projection{
			EffectiveDate: time.Date(1, time.July, 1, 0, 0, 0, 0, time.UTC),
			DebtProjections: []*debtModels.DebtProjection{
				&debtModels.DebtProjection{
					Debt:         buildDebt(0.00, 25.00),
					DebtTotal:    0.00,
					PaymentSum:   0.00,
					UnappliedSum: 25.00,
				},
				&debtModels.DebtProjection{
					Debt:         buildDebt(25.00, 25.00),
					DebtTotal:    0.00,
					PaymentSum:   25.00,
					UnappliedSum: 25.00,
				},
			},
			SavingsProjections: []*savingsModels.SavingsProjection{
				&savingsModels.SavingsProjection{
					SavingsAccount: buildSavingsAccount(1500.00, 100.00),
					SavingsTotal:   1625.00,
					PaymentSum:     125.00,
				},
			},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Savings Account Refreshes don't match. Actual: %+v; Expected %+v", actual, expected)
	}
}

func buildDebt(debtTotal float64, debtPayment float64) *debtModels.Debt {
	payments := []*debtModels.DebtPayment{
		&debtModels.DebtPayment{
			Amount:    debtPayment,
			CarryOver: true,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	return &debtModels.Debt{
		DebtName:     "Jazz 1",
		DebtTotal:    debtTotal,
		Payments:     payments,
		InterestRate: 0.00,
	}
}

func buildSavingsAccount(initialCapital float64, contributionAmount float64) *savingsModels.SavingsAccount {
	payments := []*savingsModels.SavingsPayment{
		&savingsModels.SavingsPayment{
			Amount:    contributionAmount,
			StartDate: nil,
			EndDate:   nil,
		},
	}
	return &savingsModels.SavingsAccount{
		Name:           "Jazz 1",
		APY:            0.00,
		InitialCapital: initialCapital,
		Payments:       payments,
		ProjectedDate:  nil,
	}
}
