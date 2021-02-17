package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(indexHandler)

	hf.ServeHTTP(recorder, req)
	if httpStatusCode := recorder.Code; httpStatusCode != http.StatusOK {
		t.Errorf("bad http status code: got %v expected %v",
			httpStatusCode, http.StatusOK)
	}
}
