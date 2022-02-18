package handler_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/info700700/calc-net-go/handler"
)

func TestCalc(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/api/calc", nil)
	w := httptest.NewRecorder()

	handler.Calc(w, req)

	resp := w.Result()
	bodyBytes, _ := io.ReadAll(resp.Body)
	actualBody := string(bodyBytes)

	const expactedBody = "3"

	if actualBody != expactedBody {
		t.Fatalf("Expected body: %s. Actual body: %s", expactedBody, actualBody)
	}
}
