package models

import (
	"fmt"
	"log"

	"database/sql"
	_ "github.com/lib/pq"

	// _ "github.com/jackc/pgx"
	// _ "github.com/jackc/pgx/stdlib"
)

var db *sql.DB // global variable to share it between main and the HTTP handler
var stmt *sql.Stmt

func Setup() error {
	var err error
	db, err = sql.Open("postgres", "user=root dbname=service_financial_development sslmode=require host=localhost")
	// db, err = sql.Open("pgx", "postgres://root:root@localhost:5432/service_financial_development")
	if err != nil {
		return err
	}

	stmt, err = db.Prepare("select id, currency_code, rounding from currencies")
	if err != nil {
		log.Fatal(err)
		fmt.Printf("\n")
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error on opening database connection: %s", err.Error())
		return err
	}

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
