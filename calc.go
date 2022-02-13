package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"

	interp "github.com/info700700/sum_interpreter_go"
)

func main() {
	http.HandleFunc("/", handleMain)
	addHandlerErrFunc("/api/calc", calc)

	log.Fatal(http.ListenAndServe(":80", nil))
}

//go:embed index.html
var mainPage string

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, mainPage)
}

func calc(w http.ResponseWriter, r *http.Request) error {
	expStr := r.URL.Query().Get("exp")

	result, err := interp.Exec(expStr)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(w, result)
	return err
}
