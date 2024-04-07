package db

import (
	"context"
	"fmt"
	"github.com/edgedb/edgedb-go"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Pool = connect()

func connect() *edgedb.Client {
	err := godotenv.Load()
	if err != nil {
		//log.Fatal("Error loading .env file")
	}

	edgeDBHost := os.Getenv("EDGEDB_HOST")
	edgeDBPort := os.Getenv("EDGEDB_PORT")
	edgeDBUser := os.Getenv("EDGEDB_USER")
	edgeDBPasswd := os.Getenv("EDGEDB_PASSWD")
	edgeDBName := os.Getenv("EDGEDB_DBNAME")

	ctx := context.Background()
	opts := edgedb.Options{
		Concurrency: 0,
	}
	pool, err := edgedb.CreateClientDSN(ctx, fmt.Sprintf("edgedb://%s:%s@%s:%s/%s", edgeDBUser, edgeDBPasswd, edgeDBHost, edgeDBPort, edgeDBName), opts)

	if err != nil {
		log.Fatal(err)
	}

	return pool
}
