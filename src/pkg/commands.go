package common

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
