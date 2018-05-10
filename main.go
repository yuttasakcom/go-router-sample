package main

import (
	"log"
	"net/http"

	"github.com/yuttasakcom/go-router-sample/routes"
)

func main() {
	r := routes.Router()

	log.Fatal(http.ListenAndServe(":3000", r))
}
