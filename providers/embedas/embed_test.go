package embedas

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEmbedSendsAPIKeyHeader(t *testing.T) {
	var gotKey string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotKey = r.Header.Get("X-API-Key")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"html":"<iframe></iframe>"}`))
	}))
	defer srv.Close()

	cl := NewClientWithHost(srv.URL, "secret-key")
	if _, err := cl.EmbedOne(Options{}, "https://example.com/x"); err != nil {
		t.Fatalf("EmbedOne: %s", err)
	}
	if gotKey != "secret-key" {
		t.Errorf("X-API-Key = %q, want %q", gotKey, "secret-key")
	}
}

func TestEmbedOmitsHeaderWhenNoKey(t *testing.T) {
	var hadHeader bool
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, hadHeader = r.Header["X-Api-Key"]
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"html":"<iframe></iframe>"}`))
	}))
	defer srv.Close()

	cl := NewClientWithHost(srv.URL, "")
	if _, err := cl.EmbedOne(Options{}, "https://example.com/x"); err != nil {
		t.Fatalf("EmbedOne: %s", err)
	}
	if hadHeader {
		t.Error("X-API-Key header should be absent when no key is set")
	}
}
