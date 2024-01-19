package repository

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type repositoryMock struct{}

func (m *repositoryMock) getAllTask() (map[string]interface{}, error) {
	// You can implement a mock implementation if needed for specific test cases
	return nil, nil
}

func TestGetAllTask(t *testing.T) {
	// Create a sample JSON response for testing
	mockResponse := `{"key": "value"}`

	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer mockServer.Close()

	// Set the mock server URL to the FIREBASE_URL environment variable
	os.Setenv("FIREBASE_URL", mockServer.URL)
	defer os.Unsetenv("FIREBASE_URL")

	// Create an instance of repositoryImpl
	repo := &repositoryImpl{}

	// Call the getAllTask method
	result, err := repo.getAllTask()

	// Check if there is an error
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if the result is not nil
	if result == nil {
		t.Fatal("Expected non-nil result, got nil")
	}

	// Check if the result matches the mock response
	expectedResult := map[string]interface{}{"key": "value"}
	if !jsonEqual(result, expectedResult) {
		t.Fatalf("Expected result %v, got %v", expectedResult, result)
	}
}

// Helper function to compare two JSON objects
func jsonEqual(a, b map[string]interface{}) bool {
	aBytes, err := json.Marshal(a)
	if err != nil {
		return false
	}

	bBytes, err := json.Marshal(b)
	if err != nil {
		return false
	}

	return string(aBytes) == string(bBytes)
}
