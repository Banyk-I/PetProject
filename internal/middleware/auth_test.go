package middleware

import (
	"github.com/spf13/viper"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthorizationMiddleware(t *testing.T) {
	viper.Set("secretKey", "mySecretKey")

	handler := AuthorizationMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer JWT_TOKEN")

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}
}
