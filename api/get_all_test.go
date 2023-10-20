package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetByID(t *testing.T) {
	// Create a test server that checks the parameters and responds with a simple JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedID := "123"

		if id := r.URL.Query().Get("ID"); id != expectedID {
			t.Fatalf("Expected ID parameter to be %s but got %s", expectedID, id)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"result": {"ID": "123", "name": "Test User"}}`))
	}))
	defer server.Close()

	bx24 := ExtendedBitrix24{webhookURL: server.URL}
	resp, err := bx24.getByID("crm.contact.get", 123)

	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if result, exists := resp["result"].(map[string]interface{}); !exists || result["name"] != "Test User" {
		t.Fatalf("Expected result with name 'Test User', but got: %v", resp)
	}
}
