package model

import "github.com/edgedb/edgedb-go"

type Article struct {
	Id         edgedb.UUID      `edgedb:"id" json:"id"`
	Num        int64            `edgedb:"num" json:"num"`
	Title      string           `edgedb:"title" json:"title"`
	Writer     string           `edgedb:"writer" json:"writer"`
	WriteDate  edgedb.LocalDate `edgedb:"write_date" json:"write_date" swaggertype:"string"`
	ArticleUrl string           `edgedb:"article_url" json:"article_url"`
	Content    string           `edgedb:"content" json:"content"`
	IsNotice   bool             `edgedb:"is_notice" json:"is_notice"`
	Files      []Files          `edgedb:"files" json:"files"`
}
