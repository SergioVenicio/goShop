package db

import (
	_ "github.com/lib/pq"

	"database/sql"
)

func ConnectToDatabase() *sql.DB {
	connection := "user=psql dbname=goShop password=psql host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func RunQuery(query string) (*sql.Rows, error) {
	conn := ConnectToDatabase()
	defer conn.Close()
	return conn.Query(query)
}

func Prepare(query string) (*sql.DB, *sql.Stmt, error) {
	conn := ConnectToDatabase()
	parsedQuery, err := conn.Prepare(query)
	return conn, parsedQuery, err
}
