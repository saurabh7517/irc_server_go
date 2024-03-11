package pkg

// import (
// 	"fmt"
// 	user "irc_server/src/objects"
// 	"log"
// )

// const NICK string = "NICK"
// const USER string = "USER"
// const WHOIS string = "WHOIS"
// const JOIN string = "JOIN"
// const QUIT string = "QUIT"
// const PRIVMSG string = "PRIVMSG"

// var commandArray []string = []string{"NICK", "USER", "WHOIS", "JOIN", "QUIT", "PRIVMSG"}
// var commandMap map[string]string = make(map[string]string)
// var ResponseCodes map[string]string = make(map[string]string)

// func InitAcceptableCommands() {
// 	for _, value := range commandArray {
// 		commandMap[value] = value
// 	}
// }

// func InitializeResponseCodes() map[string]string {
// 	ResponseCodes["RPL_WELCOME"] = "001"
// 	ResponseCodes["ERR_NICKNAMEINUSE"] = "433"
// 	return ResponseCodes
// }

// func isValidCommand(command *string) bool {
// 	return commandMap[*command] == *command
// }

// func ProcessCommand(parameters []string) {
// 	var command *string = &parameters[0]
// 	var response string
// 	if isValidCommand(&parameters[0]) {
// 		switch *command {
// 		case NICK:
// 			response = processNick(parameters[1:])
// 		case USER:
// 			response = processUser()
// 		case WHOIS:
// 			processWhoIs()
// 		case JOIN:
// 			processJoin()
// 		case QUIT:
// 			processQuit()
// 		case PRIVMSG:
// 			processPrivMsg()
// 		default:
// 			processDefault()

// 		}
// 	}
// }

// func processNick(nickName string) {
// 	//now check if the nickname exists
// 	//connect to database and find out
// 	//if nick name already exist return ERR_NICKNAMEINUSE
// 	// if not :bar.example.com(server name) 001(RPL_WELCOME) :<message> <nick>!<user>@<host>
// 	// rows, err := db.Query("SELECT * FROM album WHERE artist = ?", artist)
// 	ProcessNickMessage(nickName)
// 	// var db *sql.DB = GetDBConnection()
// 	rows, err := db.Query("SELECT nickname from user where nickname=? limit 1", nickName)
// 	if err != nil {
// 		log.Printf("Error retrieving rows from database")
// 	}
// 	defer rows.Close()

// 	var foundUser user.User
// 	for rows.Next() {
// 		err := rows.Scan(&foundUser.Id, &foundUser.Nickname, &foundUser.Name)
// 		if err != nil {
// 			log.Printf(err.Error())
// 			//return appropriate error to client
// 		}
// 		//user found in database
// 		//store this user in server cache
// 		//let every server know about it, and donot allow any user connection with this nickname
// 	}

// }

// func ProcessNickMessage(nickName string) string {
// 	//check if nick exists in cache
// 	//yes -> user already exists with a connection on server
// 	//no -> check if user exists in database
// 	//no -> no -> user is not registered
// 	//yes -> add user to cache send a successful response
// 	var response string = ""
// 	if IsUserOnServer(nickName) {
// 		//return message "Nickname already in use"
// 		response = fmt.Sprintf("%s %s %s:%s", ResponseCodes["ERR_NICKNAMEINUSE"], "*", nickName, "Nickname already in user")
// 	} else {
// 		if IsUserRegistered(nickName) {
// 			AddUserToCache(nickName)
// 			// send 001 RPL_WELCOME "Welcome to the Internet Relay Network<nick>!<user>@<host>"
// 			response = fmt.Sprintf("%s %s %s:%s", ResponseCodes["RPL_WELCOME"], "RPL_WELCOME", nickName, "Welcome to the Internet Relay Network")
// 		} else {
// 			//the user trying to connect is not registered, contact admin@admin.com for registration
// 			//send 502 RPL_USERNOTFOUND
// 			response = fmt.Sprintf("%s %s %s:%s", ResponseCodes["RPL_WELCOME"], "RPL_WELCOME", nickName, "Welcome to the Internet Relay Network")
// 		}
// 	}
// 	return response
// }

// func IsUserRegistered(usr string) bool {
// 	// var db *sql.DB = GetDBConnection()
// 	var found bool = false
// 	rows, err := db.Query("SELECT nickname from user where nickname=? limit 1", usr)
// 	if err != nil {
// 		log.Panic("Error retrieving rows from database")
// 	} else {
// 		var foundUser user.User
// 		for rows.Next() {
// 			err := rows.Scan(&foundUser.Id, &foundUser.Nickname, &foundUser.Name)
// 			if err != nil {
// 				log.Printf(err.Error())
// 				//return appropriate error to client
// 			}
// 			found = true
// 			break
// 		}
// 	}
// 	defer rows.Close()
// 	return found
// }

// func processUser() string {

// }

// func processWhoIs() string {

// }

// func processJoin() string {

// }

// func processQuit() string {

// }

// func processPrivMsg() string {

// }

// func processDefault() string {

// }
