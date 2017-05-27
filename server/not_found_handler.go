package server

import "net/http"

type notFoundHandler struct {
}

func (notFound *notFoundHandler) ServeHTTP(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.WriteHeader(404)
	respWriter.Write([]byte("Oops. Page not found."))
}
