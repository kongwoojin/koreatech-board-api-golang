package model

type Files struct {
	FileName string `edgedb:"file_name" json:"file_name"`
	FileUrl  string `edgedb:"file_url" json:"file_url"`
}
