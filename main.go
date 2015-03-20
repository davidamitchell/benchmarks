package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/balanceit/accounting_service/controllers"
	"github.com/balanceit/accounting_service/models"
)

var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)

func main() {

	err := models.Setup()
	if err != nil {
		fmt.Println(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Index)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}
