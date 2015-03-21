package main

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
)

type Bench struct {
	Name      string `json:"name"`
	Reference string `json:"reference"`
}

func main() {
	fmt.Println("struct to json")
	struct_to_json()

	fmt.Println("json to struct")
	json_to_struct()
}

func struct_to_json() {

	bench := Bench{"bob jones", "his reference"}

	for k := 0; k < 3; k++ {
		n := int(math.Pow(1000, float64(k)))
		start := time.Now()
		for i := 0; i < n; i++ {
			json.Marshal(bench)
		}

		elapsed := time.Since(start)
		Results(elapsed, n)
	}
}

func json_to_struct() {

	j := []byte("{\"name\":\"bob jones\",\"reference\":\"his reference\"}")

	bench := &Bench{}

	for k := 0; k < 3; k++ {
		n := int(math.Pow(1000, float64(k)))

		start := time.Now()
		for i := 0; i < n; i++ {
			json.Unmarshal(j, &bench)
		}
		elapsed := time.Since(start)
		Results(elapsed, n)
	}

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
