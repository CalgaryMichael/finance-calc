package savings

type SavingsProjection struct {
	Id             int `json:"id"`
	SavingsAccount *SavingsAccount
	SavingsTotal   float64 `json:"savingsTotal"`
	PaymentSum     float64 `json:"paymentSum"`
}
