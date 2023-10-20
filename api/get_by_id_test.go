package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetByID tests the getByID method
func TestGetAll(t *testing.T) {
	callCount := 0

	// Mocked data for pagination
	data := []map[string]interface{}{
		{"ID": "1", "name": "Test User 1"},
		{"ID": "2", "name": "Test User 2"},
		{"ID": "3", "name": "Test User 3"},
		{"ID": "4", "name": "Test User 4"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if callCount < 2 {
			resp := map[string]interface{}{
				"result": data[callCount*2 : callCount*2+2],
				"next":   float64(callCount + 1),
			}
			jsonResp, _ := json.Marshal(resp)
			w.Write(jsonResp)
			callCount++
		} else {
			resp := map[string]interface{}{
				"result": data[callCount*2:],
			}
			jsonResp, _ := json.Marshal(resp)
			w.Write(jsonResp)
		}
	}))
	defer server.Close()

	bx24 := ExtendedBitrix24{webhookURL: server.URL}
	resp, err := bx24.getAll("crm.contact.list")

	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if len(resp) != len(data) {
		t.Fatalf("Expected %d results, but got: %d", len(data), len(resp))
	}

	for i, item := range resp {
		if item["name"] != data[i]["name"] {
			t.Fatalf("Expected name %s but got %s", data[i]["name"], item["name"])
		}
	}
}
