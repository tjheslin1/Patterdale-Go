package server

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tjheslin1/Patterdale/testutil"
)

func TestCloseHandler_ServeHTTP(t *testing.T) {
	respRecorder := httptest.NewRecorder()

	quit := make(chan bool, 1)
	testLogger := testutil.NewTestLogger()

	closeHandler := &closeHandler{quit, testLogger.Logger}
	closeHandler.ServeHTTP(respRecorder, nil)

	var expectedLogMessage = "Closing server."
	if !strings.Contains(testLogger.LogOutput(), expectedLogMessage) {
		t.Errorf("Expected log output:\n'%s'\nbut was:\n'%s'\n", expectedLogMessage, testLogger.LogOutput())
	}

	closeMsg := <-quit
	if closeMsg != true {
		t.Errorf("Expected message 'true' on close channel but was '%v'\n", closeMsg)
	}
}
