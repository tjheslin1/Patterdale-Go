package database

import (
	"database/sql"
	"fmt"
	"log"
)

// DBClient represents a connection to a database.
type DBClient interface {
	Connect(dataSourceName string, logger *log.Logger)
	HealthCheck() bool
	Close()
}

// OracleDBClient represents a connection to an Oracle database.
type OracleDBClient struct {
	Logger   *log.Logger
	dbHandle *sql.DB
}

// Connect opens a connection to the Oracle database, based on the provided
// connection string.
func (dbClient *OracleDBClient) Connect(dataSourceName string, logger *log.Logger) {
	db, err := sql.Open("oci8", dataSourceName)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = db.Ping(); err != nil {
		logger.Printf("Error connecting to the database: %s\n", err)
		return
	}
}

// HealthCheck performs a simple query against the Oracle database.
func (dbClient *OracleDBClient) HealthCheck() bool {
	rows, err := dbClient.dbHandle.Query("SELECT 2+2 FROM dual")
	if err != nil {
		dbClient.Logger.Println("Error during health check")
		dbClient.Logger.Println(err)
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var sum int
		rows.Scan(&sum)
		if sum != 4 {
			return false
		}
		return true
	}

	return false
}

// Close closes the underlying database handle. This should be called in
// a deferred manner.
func (dbClient *OracleDBClient) Close() {
	dbClient.dbHandle.Close()
}
