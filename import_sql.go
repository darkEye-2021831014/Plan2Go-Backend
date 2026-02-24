package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := os.Getenv("DB_URL")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Only run if IMPORT_SQL=true
	if os.Getenv("IMPORT_SQL") == "true" {
		sqlBytes, err := ioutil.ReadFile("dump.sql")
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(string(sqlBytes))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Database imported successfully!")
		return
	}

	fmt.Println("Server would start normally here")
	// Keep your original server code below if needed
}
