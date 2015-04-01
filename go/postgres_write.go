package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"database/sql"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("pg driver")
	testing.Benchmark(benchmark)

	fmt.Println("pg driver prepared")
	testing.Benchmark(benchmark_pg_prepared)

	fmt.Println("pgx driver")
	testing.Benchmark(benchmark_pgx)

	fmt.Println("pgx driver prepared")
	testing.Benchmark(benchmark_pgx_prepared)
}

func benchmark(t *testing.B) {
	t.StopTimer()
	db, err := sql.Open("postgres", "user=root dbname=benchmarking sslmode=disable host=localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	t.StartTimer()

	start := time.Now()
	for i := 0; i < t.N; i++ {
		_, err := db.Exec("insert into bench (name, reference) values ('name', 'reference')")
		if err != nil {
			log.Fatal(err)
		}
	}
	elapsed := time.Since(start)
	Results(elapsed, t.N)
}

func benchmark_pgx(t *testing.B) {
	t.StopTimer()
	//db, err := sql.Open("postgres", "user=root dbname=service_financial_development sslmode=require host=localhost")
	db, err := sql.Open("pgx", "postgres://root:root@localhost:5432/benchmarking")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	t.StartTimer()

	start := time.Now()
	for i := 0; i < t.N; i++ {
		_, err := db.Exec("insert into bench (name, reference) values ('name', 'reference')")
		if err != nil {
			log.Fatal(err)
		}
	}
	elapsed := time.Since(start)
	Results(elapsed, t.N)
}

func benchmark_pg_prepared(t *testing.B) {
	t.StopTimer()

	db, err := sql.Open("postgres", "user=root dbname=benchmarking sslmode=disable host=localhost")
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

	ins, err := tx.Prepare("insert into bench (name, reference) values ($1, $2)")
	if err != nil {
		log.Fatal(err)
	}
	defer ins.Close()

	start := time.Now()
	for i := 0; i < t.N; i++ {
		_, err := ins.Exec("name", "reference")
		if err != nil {
			log.Fatal(err)
		}
	}

	elapsed := time.Since(start)
	Results(elapsed, t.N)
}

func benchmark_pgx_prepared(t *testing.B) {
	t.StopTimer()

	db, err := sql.Open("pgx", "postgres://root:root@localhost:5432/benchmarking")
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

	ins, err := tx.Prepare("insert into bench (name, reference) values ($1, $2)")
	if err != nil {
		log.Fatal(err)
	}
	defer ins.Close()

	start := time.Now()
	for i := 0; i < t.N; i++ {
		_, err := ins.Exec("name", "reference")
		if err != nil {
			log.Fatal(err)
		}
	}

	elapsed := time.Since(start)
	Results(elapsed, t.N)
}

func Results(elapsed time.Duration, n int) {
	each := elapsed / time.Duration(n)
	fmt.Printf(" %13.5f ms per query   | %9d iterations    | %8.5f ms total time\n", milliseconds(each), n, milliseconds(elapsed))
}

func milliseconds(dur time.Duration) float64 {
	nanoseconds := float64(dur.Nanoseconds())
	milliseconds := nanoseconds / 1000000.0
	return milliseconds
}
