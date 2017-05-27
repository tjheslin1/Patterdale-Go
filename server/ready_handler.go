package server

import "net/http"

type readyHandler struct {
}

func (readyHandler *readyHandler) ServeHTTP(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.WriteHeader(204)
}
