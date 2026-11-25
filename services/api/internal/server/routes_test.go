package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	// Arrange: creating a fake request
	req1 := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec1 := httptest.NewRecorder()

	req2 := httptest.NewRequest(http.MethodGet, "/random", nil)
	rec2 := httptest.NewRecorder()

	//Act: execute the handler
	NewRouter().ServeHTTP(rec1, req1)
	NewRouter().ServeHTTP(rec2, req2)

	//Assert 1: /healthz should give 200 status 0K
	if rec1.Code != http.StatusOK {
		t.Fatalf("status code = %d, want %d", rec1.Code, http.StatusOK)
	}

	//Assert 2: /random should give 404 Not found
	if rec2.Code != http.StatusNotFound {
		t.Fatalf("status code = %d, want %d", rec2.Code, http.StatusNotFound)
	}

	//Assert 3: the response body should be {status: ok}

	const want = `{"status": "ok"}`

	if got := rec1.Body.String(); got != want {
		t.Errorf("body = %q, want %q", got, want)
	}

}
