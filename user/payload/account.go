package payload

type AccountInfo struct {
	Id        string         `json:"id"`
	Email     string         `json:"email"`
	Invoices  []*InvoiceInfo `json:"invoices"`
	CreatedAt int64          `json:"created_at"`
}
