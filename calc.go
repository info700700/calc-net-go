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
  <style>
    label {
      display: block;
	  margin-bottom: 0.1rem;
    }
</style>
</head>
<body>
  <label for="exp">Введите выражение</label>
  <input id="exp" autofocus>
  <button>Вычислить</button>
</body>
</html>
`

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, mainPage)
}

func handleCalc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
