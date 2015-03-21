package main

import (
	"fmt"
	"log"
	"time"

	"database/sql"
	// _ "github.com/lib/pq"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "postgres://root:root@localhost:5432/benchmarking")
	//db, err := sql.Open("postgres", "user=root dbname=benchmarking sslmode=disable host=localhost")

	start := time.Now()

	rows, err := db.Query("select name from bench")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s is %d\n", name, 10)
	}

	elapsed := time.Since(start)
	log.Printf("Time elapsed %s", elapsed)
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

// func main() {
// 	var (
// 		name      string
// 		reference string
// 	)
// 	conn, err := sql.Open("pgx", "postgres://root:root@localhost:5432/benchmarking")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()
// 	start := time.Now()
// 	// Send the query to the server. The returned rows MUST be closed
// 	// before conn can be used again.
// 	rows, err := conn.Query("SELECT name, reference from bench")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	// rows.Close is called by rows.Next when all rows are read
// 	// or an error occurs in Next or Scan. So it may optionally be
// 	// omitted if nothing in the rows.Next loop can panic. It is
// 	// safe to close rows multiple times.
// 	defer rows.Close()
//
// 	// Iterate through the result set
// 	for rows.Next() {
// 		err = rows.Scan(&name, &reference)
// 		// fmt.Println(name, reference)
// 	}
//
// 	elapsed := time.Since(start)
// 	fmt.Printf("Time elapsed %s \n", elapsed)
//
// 	// Any errors encountered by rows.Next or rows.Scan will be returned here
// 	if rows.Err() != nil {
// 		log.Fatal(err)
// 	}
// }

// func main2() {
//
// 	tx, err := conn.Begin()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer tx.Commit()
//
// 	sel, err := tx.Prepare("SELECT name, reference from bench")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer sel.Close()
//
// 	start := time.Now()
//
// 	for i := 0; i < 100; i++ {
// 		rows, err := sel.Query()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
//
// 		for rows.Next() {
// 			var (
// 				name      string
// 				reference string
// 			)
// 			err = rows.Scan(&name, &reference)
// 			// fmt.Println(name, reference)
// 		}
// 		rows.Close()
// 	}
//
// 	elapsed := time.Since(start)
// 	log.Printf("Time elapsed %s", elapsed)
// }
