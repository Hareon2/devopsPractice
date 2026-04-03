package cors

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDisable_PreflightOptions(t *testing.T) {
	handler := Disable(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ok")
	})

	req := httptest.NewRequest(http.MethodOptions, "/query", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	if got := w.Header().Get("Access-Control-Allow-Origin"); got != "*" {
		t.Errorf("Access-Control-Allow-Origin = %q, want *", got)
	}
	if got := w.Header().Get("Access-Control-Allow-Methods"); got != "*" {
		t.Errorf("Access-Control-Allow-Methods = %q, want *", got)
	}
	if got := w.Header().Get("Access-Control-Allow-Headers"); got != "*" {
		t.Errorf("Access-Control-Allow-Headers = %q, want *", got)
	}
}

func TestDisable_PreservesResponseBody(t *testing.T) {
	const body = "hello-world"
	handler := Disable(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, body)
	})

	req := httptest.NewRequest(http.MethodPost, "/query", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	got := w.Body.String()
	if got != body {
		t.Errorf("body = %q, want %q", got, body)
	}
}

func TestDisable_PreservesStatusCode(t *testing.T) {
	handler := Disable(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("status = %d, want %d", w.Code, http.StatusCreated)
	}
}
