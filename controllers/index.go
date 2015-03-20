package controllers

import (
	"net/http"

	"github.com/balanceit/accounting_service/middleware"
	"github.com/balanceit/accounting_service/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		currencies := models.GetCurrencies()
		middleware.RespondJson(w, currencies, http.StatusOK)
	}
}
