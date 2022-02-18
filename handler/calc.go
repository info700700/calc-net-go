package handler

import (
	"fmt"
	"net/http"

	interp "github.com/info700700/sum_interpreter_go"
)

func Calc(w http.ResponseWriter, r *http.Request) error {
	expStr := r.URL.Query().Get("exp")

	result, err := interp.Exec(expStr)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(w, result)
	return err
}
