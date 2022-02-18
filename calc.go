package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"github.com/info700700/calc-net-go/handler"
)

func main() {
	http.HandleFunc("/", handleMain)
	addHandlerErrFunc("/api/calc", handler.Calc)

	log.Fatal(http.ListenAndServe(":80", nil))
}

//go:embed index.html
var mainPage string

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, mainPage)
}
