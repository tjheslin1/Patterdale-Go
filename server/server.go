package server

import (
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"

	"github.com/gorilla/mux"
)

// Port is the http port the server is started on.
// The default is 7000.
var Port = 7000

// Start starts up the http rest server.
func Start(logger *log.Logger, quit chan<- bool) {
	muxRouter := mux.NewRouter()

	muxRouter.NotFoundHandler = logRequestResponse(&notFoundHandler{}, logger)
	muxRouter.Handle("/ready", logRequestResponse(&readyHandler{}, logger)).Methods("GET")
	muxRouter.Handle("/close", logRequestResponse(&closeHandler{quit, logger}, logger)).Methods("POST")

	startServer(muxRouter, logger)
}

func logRequestResponse(handler http.Handler, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(respWriter http.ResponseWriter, req *http.Request) {
		loggingRespWriter := &loggingResponseWriter{
			ResponseWriter: respWriter,
		}
		requestDump, err := httputil.DumpRequest(req, true)
		check(err, logger)

		logger.Printf("REQUEST:\n%v\n::::::\n", string(requestDump))
		handler.ServeHTTP(loggingRespWriter, req)
		logger.Printf("RESPONSE:\n%d\n%s\n::::::\n", loggingRespWriter.status, string(loggingRespWriter.body))
	})
}

// startServer sets up the HTTP server in a goroutine and waits for it to exit
func startServer(handler http.Handler, logger *log.Logger) {
	go func() {
		err := http.ListenAndServe(":"+strconv.Itoa(Port), handler)
		if err != nil {
			logger.Println(err)
			panic(err)
		}
	}()

	logger.Printf("Server started on port: %v\n", strconv.Itoa(Port))
}

func check(err error, logger *log.Logger) {
	if err != nil {
		logger.Printf("Error occured handling request:\n'%v'", err)
	}
}
