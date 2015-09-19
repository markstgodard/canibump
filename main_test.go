package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestServer(resp string) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, resp)
	}))
	return ts
}

func TestCheck(t *testing.T) {
	ts := newTestServer(`{"can_i_bump": true }`)
	defer ts.Close()

	if ok := check(ts.URL); !ok {
		t.Fatalf("Expected canibump to be 'true'")
	}
}

func TestCheckCantBump(t *testing.T) {
	ts := newTestServer(`{"can_i_bump": false }`)
	defer ts.Close()

	if ok := check(ts.URL); ok {
		t.Fatalf("Expected canibump to be 'false'")
	}
}
