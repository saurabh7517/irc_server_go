package common

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initializeDB() {
	var cfg mysql.Config = mysql.Config{
		User:   "",
		Passwd: "",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "mydb",
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	log.Println("Database Connected")
}

func GetDBConnection() *sql.DB {
	return db
	// if (db!=nil) {
	// 	return db, nil;
	// }
	// return nil, errors.New("Database not initialized")
}
