package main

import (
	_ "embed"
	"fmt"
	"html"
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
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
