package payload

type InvoiceInfo struct {
	Id        string  `json:"id"`
	Email     string  `json:"email"`
	Amount    float32 `json:"amount"`
	CreatedAt int64   `json:"created_at"`
}
