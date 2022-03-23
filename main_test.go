package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health", bytes.NewBufferString(""))

	healthHandler(wr, req)

	if got := wr.Code; got != http.StatusOK {
		t.Errorf("Got %v wanted %v", got, http.StatusOK)
	}

	out, err := ioutil.ReadAll(wr.Body)
	if err != nil {
		t.Errorf("Unable to read body: %v", err)
	}

	want := "HEALTHY"
	if string(out) != want {
		t.Errorf("Got %v wanted %v", string(out), want)
	}
}
