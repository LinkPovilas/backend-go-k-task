package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LinkPovilas/backend-go-k-task/models"
)

func TestPingRoute(t *testing.T) {
	r := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("expected status code 200, got %d", w.Code)
	}

	body, err := io.ReadAll(w.Body)
	if err != nil {
		t.Errorf("error reading response body: %v", err)
	}

	var health models.Health
	if err = json.Unmarshal(body, &health); err != nil {
		t.Errorf("error unmarshalling response body: %v", err)
	}

	if health.Status != "OK" {
		t.Errorf("expected status OK, got %s", health.Status)
	}
}
