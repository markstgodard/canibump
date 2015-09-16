package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheck(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"can_i_bump": true }`)
	}))
	defer ts.Close()

	ok := check(ts.URL)
	if !ok {
		t.Fail()
	}

}

func TestCheckFail(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"can_i_bump": false }`)
	}))
	defer ts.Close()

	ok := check(ts.URL)
	if ok {
		t.Fail()
	}
}
