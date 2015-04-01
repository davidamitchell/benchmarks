package main

import (
	"fmt"
	"log"
  "math/rand"

	"database/sql"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
)

func main() {
	fmt.Println("throwing a shit tin (1 million rows) of data into the db")
	insert_many(1000000)
}

func insert_many(c int) {

	db, err := sql.Open("pgx", "postgres://root:root@localhost:5432/benchmarking")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

	for i := 0; i < c; i++ {
		_, err := ins.Exec(randSeq(20), randSeq(50))
		if err != nil {
			log.Fatal(err)
		}
	}
}
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}
