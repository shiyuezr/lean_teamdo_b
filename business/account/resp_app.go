package account

type RUser struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"Token"`
}
