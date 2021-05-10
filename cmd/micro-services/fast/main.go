package main

import (
	"github.com/gorilla/mux"
	"github.com/poncheska/k8s-microservice/pkg/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	cfgVal := os.Getenv("CFG_VAL")

	port := "8080"

	fastHandler := &handlers.FastHandler{
		ConfigValue: cfgVal,
	}
	r := mux.NewRouter()
	r.Handle("/", fastHandler)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
