package server

import (
	"log"
	"net/http"
)

type closeHandler struct {
	quit   chan<- bool
	logger *log.Logger
}

func (closeHandler *closeHandler) ServeHTTP(respWriter http.ResponseWriter, req *http.Request) {
	closeHandler.logger.Println("Closing server.")
	closeHandler.quit <- true
}
