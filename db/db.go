package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	url := os.Getenv("DATABASE_URL")

	var err error
	DB, err = sql.Open("postgres", url)

	if err != nil {
		log.Fatal("error opening connection: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("error pinging DB: ", err)
	}

	if err == nil {
		fmt.Println("Connected to DB")
	}

	CreateStockTable()
}
