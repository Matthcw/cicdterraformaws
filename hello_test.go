package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// comment 1
func TestHelloHandler(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloHandler)

	// Call the handler function directly and pass in our Request and ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := Response{Message: "Hello, Go Microservice!"}
	var got Response
	err = json.Unmarshal(rr.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	if got != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", got, expected)
	}
}
