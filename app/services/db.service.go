package services

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ilies-a/artom-api/app/utils"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	psqlconn := os.Getenv("DB_CONNECTION_URL")

	// open database
	var err error = nil
	DB, err = sql.Open("postgres", psqlconn)
	utils.CheckError(err)

	// check DB
	err = DB.Ping()
	utils.CheckError(err)

	fmt.Println("Connected to DB")
}

func CloseDB() {
	DB.Close()
}
