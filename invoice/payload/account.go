package payload

type AccountInfo struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
}
