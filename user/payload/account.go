package payload

type AccountInfo struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt int    `json:"created_at"`
}
