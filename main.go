package main

import (
	"log"
	"os"

	"github.com/tjheslin1/Patterdale/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Patterdale!")

	quit := make(chan bool)

	go server.Start(logger, quit)

	<-quit
}
