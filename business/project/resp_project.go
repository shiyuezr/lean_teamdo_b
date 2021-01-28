package project

type RProject struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Status  int    `json:"status"`
}
