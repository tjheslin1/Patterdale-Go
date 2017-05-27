package server

import (
	"net/http/httptest"
	"testing"
)

func TestReadyHandler_ServeHTTP(t *testing.T) {
	respRecorder := httptest.NewRecorder()

	readyHandler := &readyHandler{}
	readyHandler.ServeHTTP(respRecorder, nil)

	var expectedCode = 204
	if respRecorder.Code != expectedCode {
		t.Errorf("Expected status code for '%d' but was '%d'", expectedCode, respRecorder.Code)
	}
}
