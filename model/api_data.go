package model

type APIData struct {
	LastPage int     `json:"last_page"`
	Posts    []Board `json:"posts"`
}
