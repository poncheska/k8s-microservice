package main

import (
	"github.com/gorilla/mux"
	"github.com/poncheska/k8s-microservice/pkg/handlers"
	"github.com/poncheska/k8s-microservice/pkg/utils/hasher"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	cfgVal := os.Getenv("CFG_VAL")
	grt, err := strconv.Atoi(os.Getenv("HANDLER_GRT"))
	if err != nil {
		log.Printf("incorrect grt env var, set default")
		grt = 1
	}
	tol, err := strconv.Atoi(os.Getenv("HANDLER_TOL"))
	if err != nil {
		log.Printf("incorrect tol env var, set default")
		tol = 1
	}

	port := "8080"

	slowHandler := &handlers.SlowHandler{
		Hr: hasher.Hasher{
			Tol:        tol,
			Goroutines: grt,
		},
		ConfigValue: cfgVal,
	}
	r := mux.NewRouter()
	r.Handle("/slow", slowHandler)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
