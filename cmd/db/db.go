package db

import (
	"context"
	"log"

	"github.com/edgedb/edgedb-go"
)

var Pool = connect()

func connect() *edgedb.Client {
	ctx := context.Background()
	opts := edgedb.Options{
		Concurrency:     4,
		CredentialsFile: "credential.json",
	}
	pool, err := edgedb.CreateClient(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	return pool
}
