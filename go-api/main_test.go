package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoot(t *testing.T) {
	server := httptest.NewServer(newMux())
	defer server.Close()

	resp, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatalf("GET / failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 got %d", resp.StatusCode)
	}
	var body map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("decode failed: %v", err)
	}
	if body["message"] != "Hello, world" {
		t.Fatalf("unexpected body: %v", body)
	}
}

func TestHealth(t *testing.T) {
	server := httptest.NewServer(newMux())
	defer server.Close()

	resp, err := http.Get(server.URL + "/health")
	if err != nil {
		t.Fatalf("GET /health failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 got %d", resp.StatusCode)
	}
	var body map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("decode failed: %v", err)
	}
	if body["status"] != "ok" {
		t.Fatalf("unexpected body: %v", body)
	}
}

func TestItem(t *testing.T) {
	server := httptest.NewServer(newMux())
	defer server.Close()

	resp, err := http.Get(server.URL + "/items/42?q=test")
	if err != nil {
		t.Fatalf("GET /items failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 got %d", resp.StatusCode)
	}
	var body map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("decode failed: %v", err)
	}
	if int(body["item_id"].(float64)) != 42 {
		t.Fatalf("unexpected item_id: %v", body["item_id"])
	}
	if body["q"] != "test" {
		t.Fatalf("unexpected q: %v", body["q"])
	}
}
