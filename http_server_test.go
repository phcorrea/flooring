package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthcheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(healthcheckHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK, "Unexpected Response's HTTP Status Code")
	assert.Equal(t, rr.Body.String(), "Floor Found!", "Unexpected Response's Body")
}
