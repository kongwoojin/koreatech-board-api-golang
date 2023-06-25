package model

type APIData struct {
	LastPage int     `json:"last_page"` // Last page of requested board
	Posts    []Board `json:"posts"`     // Data of board
}
