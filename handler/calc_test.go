package handler_test

import (
	"net/http/httptest"
	"testing"

	"github.com/info700700/calc-net-go/handler"
)

func TestCalc(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/api/calc", nil)
	w := httptest.NewRecorder()

	handler.Calc(w, req)
}
