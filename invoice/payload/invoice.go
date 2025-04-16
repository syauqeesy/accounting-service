package payload

type InvoiceInfo struct {
	Id        string  `json:"id"`
	Email     string  `json:"email"`
	Amount    float64 `json:"amount"`
	CreatedAt int     `json:"created_at"`
}
