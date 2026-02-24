package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() (*sql.DB, error) {
	// Format: username:password@tcp(host:port)/database_name
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		// fallback to local for testing
		dsn = "root:@tcp(127.0.0.1:3306)/plan2go"
	}

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error creating DB handle:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	fmt.Println("Connected to MySQL Database successfully! 🟢")
	return DB, nil
}
