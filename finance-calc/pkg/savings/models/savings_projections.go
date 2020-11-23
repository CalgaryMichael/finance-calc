package savings

type SavingsProjection struct {
	SavingsAccount *SavingsAccount
	SavingsTotal   float64 `json:"savingsTotal"`
	PaymentSum     float64 `json:"paymentSum"`
}
