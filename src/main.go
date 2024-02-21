package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	common "irc_server/src/pkg"
	"log"
	"net"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func main() {
	//initialize logger
	//server init
	common.InitAcceptableCommands()
	initializeDB()
	// initializeServer()

	// fmt.Println(readUserInput())
}

func initializeDB() {
	var db *sql.DB
	var cfg mysql.Config = mysql.Config{
		User:   "s",
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
	fmt.Println("Connected!")
}

func initializeServer() {
	log.Println("Server start up...")
	var hostAddress string = "127.0.0.1"
	var hostPort string = "8000"
	var address string = fmt.Sprint("%s:%s", hostAddress, hostPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error initializing port :: %s on host :: %s for listening to connections", hostAddress, hostPort)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("Host :: %s and Port :: %s failed to accept connection", hostAddress, hostPort)
	}
	log.Printf("Host :: %s and Port :: %s initialized for accepting connections", hostAddress, hostPort)

	for {
		dtoMsg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println("Error reading transfer message from socket")
		}
		if isMsgOfValidLength(&dtoMsg) {
			parseIncomingMsg(&dtoMsg)
		}

	}
}

func isMsgOfValidLength(msg *string) bool {
	var strLen int = len(*msg)
	if strLen <= 510 {
		return true
	} else {
		return false
	}
}

func parseIncomingMsg(msg *string) error {
	var stringArray []string = strings.Split(*msg, " ")
	if len(stringArray) <= 16 {
		return errors.New("The command string should contain at most 15 parameters, the size of the parameters is more than 15 !!")
	} else {
		common.ProcessCommand(stringArray)

	}
	return errors.New("temp error") //FIXME

}

func readUserInput() string {
	fmt.Print("> ")
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	inputString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading input form user")
	}
	return inputString
}
