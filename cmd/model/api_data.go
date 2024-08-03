package model

type APIData struct {
	StatusCode int     `json:"status_code"` // Status code of request
	LastPage   int     `json:"last_page"`   // Last page of requested board
	Posts      []Board `json:"posts"`       // Data of board
	Error      string  `json:"error"`       // Error message
}
