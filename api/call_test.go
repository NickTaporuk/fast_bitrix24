package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCallMethod(t *testing.T) {
	// Create a test server that responds with a simple JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"result": "Test passed"}`))
	}))
	defer server.Close()

	bx24 := ExtendedBitrix24{webhookURL: server.URL}
	resp, err := bx24.CallMethod("test.method", map[string]string{})

	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if result, exists := resp["result"].(string); !exists || result != "Test passed" {
		t.Fatalf("Expected 'Test passed', but got: %v", resp)
	}
}
