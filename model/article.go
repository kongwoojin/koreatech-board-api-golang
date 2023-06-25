package model

import "github.com/edgedb/edgedb-go"

type Article struct {
	Id         edgedb.UUID      `edgedb:"id" json:"id"`
	Title      string           `edgedb:"title" json:"title"`
	Writer     string           `edgedb:"writer" json:"writer"`
	WriteDate  edgedb.LocalDate `edgedb:"write_date" json:"write_date" swaggertype:"string"`
	ArticleUrl string           `edgedb:"article_url" json:"article_url"`
	Content    string           `edgedb:"content" json:"content"`
	Files      []Files          `edgedb:"files" json:"files"`
}
