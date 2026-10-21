package health

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	// Arrange: create a fake request and response recorder.
	req1 := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec1 := httptest.NewRecorder()

	req2 := httptest.NewRequest(http.MethodPost, "/healthz", nil)
	rec2 := httptest.NewRecorder()

	// Act: execute the handler.
	HealthHandler().ServeHTTP(rec1, req1)
	HealthHandler().ServeHTTP(rec2, req2)

	// Assert 1: POST should be rejected.
	if rec2.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status code = %d, want %d", rec2.Code, http.StatusMethodNotAllowed)
	}

	if got := rec2.Header().Get("Allow"); got != http.MethodGet {
		t.Errorf("Allow = % q, want %q", got, http.MethodGet)
	}

	// Assert 2: the response status should be 200 OK.
	if rec1.Code != http.StatusOK {
		t.Fatalf("status code = %d, want %d", rec1.Code, http.StatusOK)
	}

	// Assert 3: the content type should be JSON.

	if got := rec1.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want %q", got, "application/json")
	}

	// Assert 4: the response body should be {status: ok}.
	const want = `{"status": "ok"}`
	if got := rec1.Body.String(); got != want {
		t.Errorf("body = %q, want %q", got, want)
	}

}
