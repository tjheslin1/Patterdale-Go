package main

import (
	"log"
	"os"

	_ "github.com/mattn/go-oci8"

	"github.com/tjheslin1/Patterdale/database"
	"github.com/tjheslin1/Patterdale/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Patterdale!")

	oracleClient := database.OracleDBClient{Logger: logger}
	oracleClient.Connect("system/oracle@localhost:1521/xe.oracle.docker", logger)

	logger.Printf("Oracle Health check = '%v'\n", oracleClient.HealthCheck())

	quit := make(chan bool)

	go server.Start(logger, quit)

	<-quit
}
