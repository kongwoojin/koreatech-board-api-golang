package model

import "github.com/edgedb/edgedb-go"

type ApiArticle struct {
	StatusCode int              `json:"status_code"` // Status code of request
	Error      string           `json:"error"`       // Error message
	Id         edgedb.UUID      `json:"id"`
	Title      string           `json:"title"`
	Writer     string           `json:"writer"`
	WriteDate  edgedb.LocalDate `json:"write_date" swaggertype:"string"`
	ArticleUrl string           `json:"article_url"`
	Content    string           `json:"content"`
	Files      []Files          `json:"files"`
}
