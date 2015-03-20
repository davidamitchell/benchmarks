package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"database/sql"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
)

func main() {
	br := testing.Benchmark(benchmark)
	fmt.Println(br)
}

func benchmark(t *testing.B) {
	t.StopTimer()
	//db, err := sql.Open("postgres", "user=root dbname=service_financial_development sslmode=require host=localhost")
	db, err := sql.Open("pgx", "postgres://root:root@localhost:5432/service_financial_development")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	t.StartTimer()

	for i := 0; i < t.N; i++ {
		rows, err := db.Query("SELECT id, currency_code, rounding from currencies")
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			var (
				id            string
				currency_code string
				rounding      string
			)
			err = rows.Scan(&id, &currency_code, &rounding)
			//fmt.Println(id, currency_code, rounding)
		}
		rows.Close()
	}
}

func benchmark(t *testing.B) {
	t.StopTimer()
	//db, err := sql.Open("postgres", "user=root dbname=service_financial_development sslmode=require host=localhost")
	db, err := sql.Open("pgx", "postgres://root:root@localhost:5432/service_financial_development")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	t.StartTimer()

	for i := 0; i < t.N; i++ {
		rows, err := db.Query("SELECT id, currency_code, rounding from currencies")
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			var (
				id            string
				currency_code string
				rounding      string
			)
			err = rows.Scan(&id, &currency_code, &rounding)
			//fmt.Println(id, currency_code, rounding)
		}
		rows.Close()
	}
}

func benchmarkPrepared(t *testing.B) {
	t.StopTimer()
	//db, err := sql.Open("postgres", "user=root dbname=service_financial_development sslmode=require host=localhost")
	db, err := sql.Open("pgx", "postgres://root:root@localhost:5432/service_financial_development")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	t.StartTimer()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()

	sel, err := tx.Prepare("SELECT id, currency_code, rounding from currencies")
	if err != nil {
		log.Fatal(err)
	}
	defer sel.Close()

	for i := 0; i < 100; i++ {
		start := time.Now()
		rows, err := sel.Query()
		if err != nil {
			log.Fatal(err)
		}
		elapsed := time.Since(start)
		log.Printf("running query %s", elapsed)

		for rows.Next() {
			var (
				id            string
				currency_code string
				rounding      string
			)
			err = rows.Scan(&id, &currency_code, &rounding)
			fmt.Println(id, currency_code, rounding)
		}
		rows.Close()
	}
}

func benchmarkOnTempTable(t *testing.B) {
	t.StopTimer()
	db, err := sql.Open("postgres", "sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	t.StartTimer()

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Commit()

	_, err = tx.Exec("CREATE TEMPORARY TABLE users(user_name TEXT, first_name TEXT, last_name TEXT)")
	if err != nil {
		t.Fatal(err)
	}
	ins, err := tx.Prepare("INSERT INTO users (user_name, first_name, last_name) VALUES('go', 'is', 'it')")
	if err != nil {
		t.Fatal(err)
	}
	defer ins.Close()
	sel, err := tx.Prepare("SELECT user_name, first_name, last_name from users LIMIT 50")
	if err != nil {
		t.Fatal(err)
	}
	defer sel.Close()
	for i := 0; i < 100000; i++ {
		if i%2 == 0 {
			_, err := ins.Exec()
			if err != nil {
				t.Fatal(err)
			}
		} else {
			rows, err := sel.Query()
			if err != nil {
				t.Fatal(err)
			}

			for rows.Next() {
				if i == 3 {
					var user_name string
					var first_name string
					var last_name string
					err = rows.Scan(&user_name, &first_name, &last_name)
					fmt.Println(user_name, first_name, last_name)
				}
			}
			rows.Close()
		}
	}
}
