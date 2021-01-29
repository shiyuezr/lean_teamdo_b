package project

import "teamdo/business/account"

type RProject struct {
	Id             int              `json:"id"`
	Name           string           `json:"name"`
	Content        string           `json:"content"`
	Status         int              `json:"status"`
	Administrators []*account.RUser `json:"administrators"`
	Participants   []*account.RUser `json:"participants"`
}
