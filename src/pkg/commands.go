package common

import (
	user "irc_server/src/objects"
	"log"
)

const NICK string = "NICK"
const USER string = "USER"
const WHOIS string = "WHOIS"
const JOIN string = "JOIN"
const QUIT string = "QUIT"
const PRIVMSG string = "PRIVMSG"

var commandArray []string = []string{"NICK", "USER", "WHOIS", "JOIN", "QUIT", "PRIVMSG"}
var commandMap map[string]string = make(map[string]string)

func InitAcceptableCommands() {
	for _, value := range commandArray {
		commandMap[value] = value
	}
}

func isValidCommand(command *string) bool {
	return commandMap[*command] == *command
}

func ProcessCommand(parameters []string) {
	var command *string = &parameters[0]
	if isValidCommand(&parameters[0]) {
		switch *command {
		case NICK:
			processNick(parameters[1:])
		case USER:
			processUser()
		case WHOIS:
			processWhoIs()
		case JOIN:
			processJoin()
		case QUIT:
			processQuit()
		case PRIVMSG:
			processPrivMsg()
		default:
			processDefault()

		}
	}
}

func processNick(parameters []string) {
	//now check if the nickname exists
	//connect to database and find out
	//if nick name already exist return ERR_NICKNAMEINUSE
	// if not :bar.example.com(server name) 001(RPL_WELCOME) :<message> <nick>!<user>@<host>
	// rows, err := db.Query("SELECT * FROM album WHERE artist = ?", artist)
	ProcessNickMessage()
	// var db *sql.DB = GetDBConnection()
	rows, err := db.Query("SELECT nickname from user where nickname=? limit 1", parameters[1])
	if err != nil {
		log.Printf("Error retrieving rows from database")
	}
	defer rows.Close()

	var foundUser user.User
	for rows.Next() {
		err := rows.Scan(&foundUser.Id, &foundUser.Nickname, &foundUser.Name)
		if err != nil {
			log.Printf(err.Error())
			//return appropriate error to client
		}
		//user found in database
		//store this user in server cache
		//let every server know about it, and donot allow any user connection with this nickname
	}

}

func ProcessNickMessage(parameters []string) {
	//check if nick exists in cache
	//yes -> user already exists with a connection on server
	//no -> check if user exists in database
	//no -> no -> user is not registered
	//yes -> add user to cache send a successful response
	if IsUserOnServer(parameters[1]) {
		//return message "Nickname already in use"
	} else {
		if IsUserRegistered(parameters[1]) {
			//add user to cache
			AddLogedInUser(parameters[1])
			// send 001 RPL_WELCOME "Welcome to the Internet Relay Network<nick>!<user>@<host>"
		} else {
			//the user trying to connect is not registered, contact admin@admin.com for registration
		}
	}
}

func IsUserRegistered(usr string) bool {
	// var db *sql.DB = GetDBConnection()
	rows, err := db.Query("SELECT nickname from user where nickname=? limit 1", usr)
	if err != nil {
		log.Printf("Error retrieving rows from database")
	}
	defer rows.Close()

	var found bool = false

	var foundUser user.User
	for rows.Next() {
		err := rows.Scan(&foundUser.Id, &foundUser.Nickname, &foundUser.Name)
		if err != nil {
			log.Printf(err.Error())
			//return appropriate error to client
		}
		found = true
		break
	}
	return found
}

func processUser() {

}

func processWhoIs() {

}

func processJoin() {

}

func processQuit() {

}

func processPrivMsg() {

}

func processDefault() {

}
