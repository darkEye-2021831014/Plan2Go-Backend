package importsql

import (
	"io/ioutil"
	"log"

	db "plan2go-backend/infra/DB"

	_ "github.com/go-sql-driver/mysql"
)

func Import() {
	dbcn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	sqlBytes, err := ioutil.ReadFile("dump.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = dbcn.Exec(string(sqlBytes))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database imported successfully!")
}
