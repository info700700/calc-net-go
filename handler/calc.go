package handler

import (
	_ "embed"
	"fmt"
	"net/http"

	interp "github.com/info700700/sum_interpreter_go"
)

//go:embed index.html
var mainPage string

func Main(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, mainPage)
}

func Calc(w http.ResponseWriter, r *http.Request) {
	expStr := r.URL.Query().Get("exp")

	result, err := interp.Exec(expStr)
	if err != nil {
		outputInternalServerError(w)
		return
	}

	_, err = fmt.Fprint(w, result)
	if err != nil {
		outputInternalServerError(w)
		return
	}
}
