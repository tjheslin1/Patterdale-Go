package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tjheslin1/GoSchedule/testutil"
)

func TestServer_startServer(t *testing.T) {
	testLogger := testutil.NewTestLogger()
	handler := logRequestResponse(&dummyHandler{}, testLogger.Logger)

	startServer(handler, testLogger.Logger)

	expectedLogOutput := "Server started on port: 7000\n"
	if testLogger.LogOutput() != expectedLogOutput {
		t.Errorf("Expected log output to be:\n'%s'\nbut was:\n'%s'\n", expectedLogOutput, testLogger.LogOutput())
	}
}

func TestServer_logRequestResponse(t *testing.T) {
	testLogger := testutil.NewTestLogger()
	handler := logRequestResponse(&dummyHandler{}, testLogger.Logger)

	req, err := http.NewRequest("GET", "/test", strings.NewReader(`{"test": "json"}`))
	if err != nil {
		t.Errorf("Error occured creating test request.\n%v\n", err)
	}

	respRecorder := httptest.NewRecorder()

	handler.ServeHTTP(respRecorder, req)

	expectedLogOutput := `REQUEST:
GET /test HTTP/1.1

{"test": "json"}
::::::
RESPONSE:
201

::::::`

	if !strings.Contains(testLogger.LogOutput(), "REQUEST") && !strings.Contains(testLogger.LogOutput(), "RESPONSE") {
		t.Errorf("Expected log output to be:\n'%s'\nbut was:\n'%s'\n", expectedLogOutput, testLogger.LogOutput())
	}
}

func TestServer_check(t *testing.T) {
	testLogger := testutil.NewTestLogger()
	err := errors.New("expected err occured")

	check(err, testLogger.Logger)

	expectedLogOutput := "Error occured handling request:\n'expected err occured'\n"
	if testLogger.LogOutput() != expectedLogOutput {
		t.Errorf("Expected log output to be:\n'%s'\nbut was:\n'%s'\n", expectedLogOutput, testLogger.LogOutput())
	}
}

type dummyHandler struct {
}

func (dummy *dummyHandler) ServeHTTP(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.WriteHeader(201)
}
