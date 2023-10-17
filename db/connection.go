package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dbUser := "root"
	dbPass := "root"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "deres"
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	database, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// Check that the database is available and accessible
	err = database.Ping()
	if err != nil {
		return nil, err
	}
	return database, nil
}
