package server

import (
	"net/http/httptest"
	"testing"
)

func TestNotFoundHandler_ServerHTTP(t *testing.T) {
	respRecorder := httptest.NewRecorder()

	notFoundHandler := &notFoundHandler{}
	notFoundHandler.ServeHTTP(respRecorder, nil)

	var expectedCode = 404
	if respRecorder.Code != expectedCode {
		t.Errorf("Expected status code for '%d' but was '%d'", expectedCode, respRecorder.Code)
	}

	var expectedBody = "Oops. Page not found."
	if string(respRecorder.Body.Bytes()) != expectedBody {
		t.Errorf("Expected response body of\n'%s'\nbut got:\n'%s\n", expectedBody, string(respRecorder.Body.Bytes()))
	}
}
