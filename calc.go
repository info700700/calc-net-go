package main

import (
	"log"
	"net/http"

	"github.com/info700700/calc-net-go/handler"
)

func main() {
	http.HandleFunc("/", handler.Main)
	http.HandleFunc("/api/calc", handler.Calc)

	log.Fatal(http.ListenAndServe(":80", nil))
}
