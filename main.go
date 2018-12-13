package main

import (
	"log"
	"net/http"

	"github.com/cosaques/fizzbuzz/handlers"
)

func main() {
	http.HandleFunc("/", handlers.FizzBuzzHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
