package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/info700700/calc-net-go/handler"
)

func TestCalc(t *testing.T) {
	// "+" кодируется как "%2B".
	req := httptest.NewRequest(http.MethodGet, "http://example.com/api/calc?exp=1%2B2", nil)
	w := httptest.NewRecorder()

	handler.Calc(w, req)

	resp := w.Result()

	{
		actualStatusCode := resp.StatusCode
		const expectedStatusCode = http.StatusOK

		if actualStatusCode != expectedStatusCode {
			t.Fatalf("Actual status code: %d. Expected status code: %d.", actualStatusCode, expectedStatusCode)
		}
	}
	{
		bodyBytes, _ := io.ReadAll(resp.Body)
		actualBody := string(bodyBytes)
		const expectedBody = "3"

		if actualBody != expectedBody {
			t.Fatalf("Actual body: %s. Expected body: %s", actualBody, expectedBody)
		}
	}
}
