package main

import "net/http"

type HandlerErrFunc func(http.ResponseWriter, *http.Request) error

func addHandlerErrFunc(pattern string, handlerErrFunc HandlerErrFunc) {
	handlerFunc := convertHandlerErrFuncToHandlerFunc(handlerErrFunc)
	http.HandleFunc(pattern, handlerFunc)
}

func convertHandlerErrFuncToHandlerFunc(f HandlerErrFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			outputInternalServerError(w)
		}
	}
}

func outputInternalServerError(w http.ResponseWriter) {
	statusCode := http.StatusInternalServerError
	statusText := http.StatusText(statusCode)
	http.Error(w, statusText, statusCode)
}
