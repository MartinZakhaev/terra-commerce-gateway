package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MartinZakhaev/terra-commerce-gateway/internal/config"
)

func TestHealthEndpoints(t *testing.T) {
	server := New(config.Config{
		ListenAddress:   ":0",
		RequestTimeout:  time.Second,
		ShutdownTimeout: time.Second,
	})

	for _, path := range []string{"/health/live", "/health/ready"} {
		t.Run(path, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, path, nil)
			response := httptest.NewRecorder()

			server.HTTP().Handler.ServeHTTP(response, request)

			if response.Code != http.StatusOK {
				t.Fatalf("expected %d, got %d", http.StatusOK, response.Code)
			}
		})
	}
}
