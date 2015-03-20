package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RespondJson(w http.ResponseWriter, i interface{}, code int) {
	j, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", j)
}
