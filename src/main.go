package main

import (
	"bufio"
	"fmt"
	object "irc_server/src/objects"
	"log"
	"net"
	"os"

	"google.golang.org/protobuf/proto"
)

const hostAddress string = "127.0.0.1"
const hostPort string = "8080"

func main() {
	//initialize logger
	//server init
	// common.InitAcceptableCommands()
	// common.InitializeResponseCodes()
	// initializeDB()
	initializeServer()

	// fmt.Println(readUserInput())
}

func initializeServer() {
	log.Println("Server start up...")
	var address string = fmt.Sprintf("%v:%v", hostAddress, hostPort)
	listener, err := net.Listen("tcp4", address)
	if err != nil {
		log.Fatalf("Error initializing %s", address)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("Host :: %s and Port :: %s failed to accept connection", hostAddress, hostPort)
	}
	log.Printf("Host :: %s and Port :: %s initialized for accepting connections", hostAddress, hostPort)

	var dtoMsg []byte = make([]byte, 512)

	for {
		dtoMsgLen, err := bufio.NewReader(conn).Read(dtoMsg)
		if err != nil {
			log.Println("Error reading transfer message from socket")
			break
		}
		if dtoMsgLen > 0 {
			log.Printf("Number of bytes read on socket %d", dtoMsgLen)

			var incomingMsg *object.Message = &object.Message{}

			proto.Unmarshal(dtoMsg, incomingMsg)

			var outgoingMsg *object.Message = &object.Message{}

			switch incomingMsg.Command {
			case object.Command_Ping:
				outgoingMsg.Command = object.Command_Pong
				outgoingMsg.HostAddress = &object.HostAddress{HostIp: hostAddress, HostPort: hostPort}
				break
			default:
				outgoingMsg.Command = object.Command_Unkwn
				outgoingMsg.HostAddress = &object.HostAddress{HostIp: hostAddress, HostPort: hostPort}
			}

			var msgBytes []byte = encodeMessage(outgoingMsg)

			conn.Write(msgBytes)

		}
	}
	log.Println("Shutting Down Server...")
	conn.Close()
	log.Println("Shut Down Complete...")
}

func encodeMessage(message *object.Message) []byte {
	msgBytes, err := proto.Marshal(message)
	if err != nil {
		log.Println("Error marshalling object")
	}
	return msgBytes
}

func isMsgOfValidLength(msg *string) bool {
	var strLen int = len(*msg)
	if strLen <= 510 {
		return true
	} else {
		return false
	}
}

// func parseIncomingCommand(msg *string) error {
// 	var stringArray []string = strings.Split(*msg, " ")
// 	if len(stringArray) <= 16 {
// 		return errors.New("The command string should contain at most 15 parameters, the size of the parameters is more than 15 !!")
// 	} else {
// 		common.ProcessCommand(stringArray)

// 	}
// 	return errors.New("temp error") //FIXME

// }

func readUserInput() string {
	fmt.Print("> ")
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	inputString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading input form user")
	}
	return inputString
}
