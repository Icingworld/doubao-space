package httpapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Icingworld/doubao-space/internal/database"
)

func TestHealthAndFrontendRoutes(t *testing.T) {
	db, err := database.Open(t.TempDir())
	if err != nil {
		t.Fatalf("database.Open() error = %v", err)
	}
	defer db.Close()

	router := NewRouter(db, "test", true)

	health := httptest.NewRecorder()
	router.ServeHTTP(health, httptest.NewRequest(http.MethodGet, "/api/health", nil))
	if health.Code != http.StatusOK {
		t.Fatalf("health status = %d, want %d", health.Code, http.StatusOK)
	}

	frontend := httptest.NewRecorder()
	router.ServeHTTP(frontend, httptest.NewRequest(http.MethodGet, "/dashboard", nil))
	if frontend.Code != http.StatusOK {
		t.Fatalf("frontend status = %d, want %d", frontend.Code, http.StatusOK)
	}
}
