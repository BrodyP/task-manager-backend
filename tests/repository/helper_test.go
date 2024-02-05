package repository_test

import (
	"net/http"
	"net/http/httptest"
	"os"
)

func mockingServer(mockingResponse []byte, fn func()) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockingResponse))
	}))
	defer mockServer.Close()
	os.Setenv("FIREBASE_URL", mockServer.URL)
	defer os.Unsetenv("FIREBASE_URL")

	fn()
}
