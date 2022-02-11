package main

import (
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

const mainPage = `
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>Калькулятор</title>
</head>
<body>
  <p>
    Введите выражение
  </p>
  <input autofocus>
</body>
</html>
`

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, mainPage)
}

func handleCalc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
