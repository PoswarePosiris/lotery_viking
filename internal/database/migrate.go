package database

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"
	"strings"
)

//go:embed schema.sql
var schema embed.FS

func Migrate() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, host, port))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + os.Getenv("DB_DATABASE"))
	if err != nil {
		return err
	}

	_, err = db.Exec("USE " + os.Getenv("DB_DATABASE"))
	if err != nil {
		return err
	}

	sqlFile, err := schema.ReadFile("schema.sql")
	if err != nil {
		return err
	}

	requests := string(sqlFile)

	// Split the SQL file into individual statements
	statements := strings.Split(requests, ";")

	// Execute each statement separately
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue // Skip empty statements
		}
		_, err := db.Exec(stmt)
		if err != nil {
			return err
		}
	}

	fmt.Println("Database schema created successfully")
	return nil
}

func Drop() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, host, port))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DROP DATABASE " + os.Getenv("DB_DATABASE"))
	if err != nil {
		return err
	}
	return nil
}
