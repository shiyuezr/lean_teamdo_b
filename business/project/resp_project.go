package project

type RProject struct {
	Id int	`json:"id"`
	Name 	string `json:"name"`
}

type RTunnel struct {
	Id 	int		`json:"id"`
	Title	string	`json:"title"`
	DisplayIndex	int	`json:"display_index"`
}