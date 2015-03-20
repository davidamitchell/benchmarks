package main

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

var conn *pgx.Conn

func main() {
	SetUp()

	err := ListTasks()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to list tasks: %v\n", err)
		os.Exit(1)
	}

}

func SetUp() {
	var err error
	conn, err = pgx.Connect(ExtractConfig())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
}

func ListTasks() error {
	rows, _ := conn.Query("select id, currency_code from currencies")

	for rows.Next() {
		var id string
		var description string
		err := rows.Scan(&id, &description)
		if err != nil {
			return err
		}
		fmt.Printf("%s %s\n", id, description)
	}
	err := rows.Err()
	rows.Close()
	return err
}

func ExtractConfig() pgx.ConnConfig {
	var config pgx.ConnConfig

	config.Host = "localhost"
	config.User = "root"

	config.Password = "root"
	config.Database = "service_financial_development"

	return config
}
