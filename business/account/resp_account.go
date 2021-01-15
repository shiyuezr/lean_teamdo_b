package account

import "time"

type RUser struct {
	Id        int       `json:"id"`
	UserName  string    `json:"user_name"`
	UserCode  string    `json:"user_code"`
	Age       int       `json:age`
	Password  string    `json:password`
	Level     int       `json:level` //0为管理员，1为普通成员
	IsEnabled bool      `json:"is_enabled"`
	IsDeleted bool      `json:"is_deleted"`
	CreateAt  time.Time `json:"create_at"`
}
