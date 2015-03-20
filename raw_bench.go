package main

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx"
)

var conn *pgx.Conn

func main() {
	br := testing.Benchmark(benchmark)
  nanoseconds := float64(br.T.Nanoseconds()) / float64(br.N)
  milliseconds := nanoseconds / 1000000.0

  fmt.Printf("%13.2f ns/op | %13.10f ms/op | %d Iterations\n", nanoseconds, milliseconds, br.N)
	fmt.Println(br)
}

func benchmark(t *testing.B) {
	t.StopTimer()
	SetUp()
	t.StartTimer()

	start := time.Now()
	for i := 0; i < t.N; i++ {
		ListTasks()
	}
	elapsed := time.Since(start) / time.Duration(t.N)
	fmt.Printf("running query %s\n", elapsed.Millisecond())
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
	rows, err := conn.Query("select id, currency_code from currencies")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	for rows.Next() {
		var id string
		var description string
		err := rows.Scan(&id, &description)
		if err != nil {
			return err
		}
		// fmt.Printf("%s %s\n", id, description)
	}
	err = rows.Err()
	rows.Close()
	return err
}

func ExtractConfig() pgx.ConnConfig {
	var config pgx.ConnConfig

	config.Host = "localhost"
	config.User = "root"

	config.Password = "root"
	config.Database = "bench_testing"

	return config
}
