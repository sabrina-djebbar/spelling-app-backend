package handler

import (
	"fmt"
	"net/http"
)

var HealthPath = "/"

var (
	HealthHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"ready":"ok"}`)
	})
)
