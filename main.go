package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/ibmdb/go_ibm_db"
)

func main() {
	connStr := toConnString(map[string]string{
		"HOSTNAME": "localhost",
		"DATABASE": "testdb",
		"PORT":     "50000",
		"UID":      "db2inst1",
		"PWD":      "testdb",
	}, ";")
	db, err := sql.Open("go_ibm_db", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer func() {
		db := db
		fmt.Println("Shutting down connection")
		_ = db.Close()
	}()
	fmt.Println("Successfully connected to the DB")
}
