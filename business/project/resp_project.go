package project

import (
	b_account "teamdo/business/account"
	"time"
)

type RProject struct {
	Id           int                `json:"id"`
	Name         string             `json:"name"`
	Introduction string             `json:"introduction"`
	Cover        string             `json:"cover"`
	StartTime    time.Time          `json:"start_time"`
	CreateAt     time.Time          `json:"create_at"`
	IsEnabled    bool               `json:"is_enabled"`
	IsDeleted    bool               `json:"is_deleted"`
	Users        []*b_account.RUser `json:"users"`
}
