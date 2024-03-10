package data

import (
	"database/sql"
	obj "irc_server/objects"
	"log"

	"github.com/go-sql-driver/mysql"
)

type sqlDb struct {
	db *sql.DB
}

var dbConn *sqlDb = &sqlDb{db: nil}

func initializeSQLDB() *sqlDb {
	var db *sql.DB
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
		log.Panic(err)
	} else {
		log.Println("Database Connected")
		dbConn.setConnection(db)
		return dbConn
	}
	return nil
}

func GetSQLDBConnection() *sqlDb {
	if dbConn.db != nil {
		return dbConn
	} else {
		conn := initializeSQLDB()
		if conn == nil {
			return nil
		}
		return conn
	}
}

func (dbConn *sqlDb) setConnection(conn *sql.DB) {
	dbConn.db = conn
}

// getOneUser(username string, password string) (*obj.User, error)
// getUserList() []obj.User
// isUserPresent(username string) bool
// insertUser(username string, password string) bool

func (dbConn *sqlDb) getOneUser(username string, password string) (*obj.User, error) {
	return nil, nil

}

func (dbConn *sqlDb) getUserList() []obj.User {
	return nil

}

func (dbConn *sqlDb) isUserPresent(username string) bool {
	return false

}

func (dbConn *sqlDb) insertUser(username string, password string) bool {
	return false

}
