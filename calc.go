package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/api/calc", handleCalc)

	log.Fatal(http.ListenAndServe(":80", nil))
}

//go:embed index.html
var mainPage string

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, mainPage)
}

func handleCalc(w http.ResponseWriter, r *http.Request) {
	expStr := r.URL.Query().Get("exp")
	fmt.Fprintf(w, "Exp: %q", expStr)
}
