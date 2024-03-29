package model

import "github.com/edgedb/edgedb-go"

type Board struct {
	Id        edgedb.UUID      `edgedb:"id" json:"id"`
	Num       int64            `edgedb:"num" json:"num"`
	Title     string           `edgedb:"title" json:"title"`
	Writer    string           `edgedb:"writer" json:"writer"`
	WriteDate edgedb.LocalDate `edgedb:"write_date" json:"write_date" swaggertype:"string"`
	ReadCount int64            `edgedb:"read_count" json:"read_count"`
	IsNew     bool             `edgedb:"is_new" json:"is_new"`
}
