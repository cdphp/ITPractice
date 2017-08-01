package model

// Token struct
type Token struct {
	ID        int64  `json:"id"`
	Token     string `json:"token"`
	UserID    int64  `json:"user_id"`
	Expire    int    `json:"expire"`
	Auth      string `json:"auth"`
	LogoutAt  int64  `json:"lougout_at"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// Obtian 获取token
func (t *Token) Obtian(UserID int64, Auth string, Expire int) string {
	return "123"
}
