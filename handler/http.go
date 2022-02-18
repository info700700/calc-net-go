package handler

import "net/http"

func outputInternalServerError(w http.ResponseWriter) {
	statusCode := http.StatusInternalServerError
	statusText := http.StatusText(statusCode)
	http.Error(w, statusText, statusCode)
}
