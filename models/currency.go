package models

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	// _ "github.com/jackc/pgx"
	// _ "github.com/jackc/pgx/stdlib"
)

type Currency struct {
	Id           string `json:"id"`
	CurrencyCode string `json:"currency_code"`
	Rounding     string `json:"rounding"`
}

func GetCurrencies() []Currency {

	var currencies []Currency
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
		fmt.Printf("\n")
	}

	var c Currency
	for rows.Next() {
		if err := rows.Scan(&c.Id, &c.CurrencyCode, &c.Rounding); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", c)
		currencies = append(currencies, c)
	}

	rows.Close() // important to close explictly and not wait for defer

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return currencies
}
