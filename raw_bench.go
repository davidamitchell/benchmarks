package main

import (
	"fmt"
	"os"
	"testing"
	"time"

	"database/sql"
	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

var conn *pgx.Conn
var db *sql.DB

func main() {
	SetUp()
	br := testing.Benchmark(benchmark)

	nanoseconds := float64(br.T.Nanoseconds()) / float64(br.N)
	milliseconds := nanoseconds / 1000000.0

	fmt.Printf("%13.2f ns/op | %13.10f ms/op | %d Iterations\n", nanoseconds, milliseconds, br.N)
	fmt.Println(br)

	fmt.Println("\n\n===================\n\n")
	SetUpPg()
	br = testing.Benchmark(benchmark_pg)

	nanoseconds = float64(br.T.Nanoseconds()) / float64(br.N)
	milliseconds = nanoseconds / 1000000.0

	fmt.Printf("%13.2f ns/op | %13.10f ms/op | %d Iterations\n", nanoseconds, milliseconds, br.N)
	fmt.Println(br)
}

func benchmark(t *testing.B) {
	t.StopTimer()
	t.StartTimer()

	start := time.Now()
	for i := 0; i < t.N; i++ {
		ListTasks()
	}
	elapsed := time.Since(start)
	Results(elapsed, t.N)
}

func benchmark_pg(t *testing.B) {
	t.StopTimer()
	t.StartTimer()

	start := time.Now()
	for i := 0; i < t.N; i++ {
		ListTasksPG()
	}
	elapsed := time.Since(start)
	Results(elapsed, t.N)
}

func Results(elapsed time.Duration, n int) {
	each := elapsed / time.Duration(n)
	fmt.Printf(" %13.10s per query   | %9d iterations    | %8s total time\n", each, n, elapsed)
}

func SetUp() {
	var err error
	conn, err = pgx.Connect(ExtractConfig())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
}

func SetUpPg() {
	var err error
	db, err = sql.Open("postgres", "user=root dbname=benchmarking sslmode=disable host=localhost")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
}

func ListTasks() error {
	rows, err := conn.Query("select name, reference from bench")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		var name string
		var reference string
		err := rows.Scan(&name, &reference)
		if err != nil {
			return err
		}
		// fmt.Printf("%s %s\n", id, description)
	}
	err = rows.Err()

	return err
}

func ListTasksPG() error {
	rows, err := db.Query("select name, reference from bench")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var reference string
		err := rows.Scan(&name, &reference)
		if err != nil {
			return err
		}
		// fmt.Printf("%s %s\n", id, description)
	}
	err = rows.Err()

	return err
}

func ExtractConfig() pgx.ConnConfig {
	var config pgx.ConnConfig

	config.Host = "localhost"
	config.User = "root"

	config.Password = "root"
	config.Database = "benchmarking"

	return config
}
