package main

import (
	"bufio"
	"fmt"
	data "irc_server/data"
	"irc_server/login"
	obj "irc_server/objects"
	util "irc_server/pkg"
	register "irc_server/registration"
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
	data.InitializeDatabaseService("IN_MEM")

	initializeServer()

	// fmt.Println(readUserInput())
}

func initializeServer() {

	log.Println("Server start up...")
	var conn net.Conn
	var listener net.Listener

	for {

		log.Println("Listening for connections...")
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
			dtoMsgLen, err := conn.Read(dtoMsg)

			if err != nil {
				log.Println("Error reading transfer message from socket")
				log.Println("Client closed connection")
				break
			}
			if dtoMsgLen > 0 {
				log.Printf("Number of bytes read on socket %d", dtoMsgLen)

				var incomingMsg *obj.Message = &obj.Message{}

				proto.Unmarshal(dtoMsg, incomingMsg)

				var outgoingMsg *obj.Message = &obj.Message{}
				var msgBytes []byte

				switch incomingMsg.Command {
				case obj.Command_Ping:
					outgoingMsg.Command = obj.Command_Pong
					outgoingMsg.HostAddress = &obj.HostAddress{HostIp: hostAddress, HostPort: hostPort}
					msgBytes = util.EncodeMessage(outgoingMsg)
				case obj.Command_Reg:
					log.Println(incomingMsg.HostAddress.HostIp)
					msgBytes = register.HandleReqRes(incomingMsg)
				case obj.Command_Log:
					log.Println(incomingMsg.HostAddress.HostIp)
					msgBytes = login.HandleReqRes(incomingMsg)
				default:
					outgoingMsg.Command = obj.Command_Unkwn
					outgoingMsg.HostAddress = &obj.HostAddress{HostIp: hostAddress, HostPort: hostPort}
					msgBytes = util.EncodeMessage(outgoingMsg)
				}

				x, err := conn.Write(msgBytes)
				if err != nil {
					log.Println(err)
				}
				log.Printf("Bytes written on socket connection %d \n", x)

			}
		}
		log.Println("Closing connection")
		conn.Close()
		listener.Close()
		log.Println("Connection closed")
	}

	log.Println("Shutting Down Server...") //TODO move server to thread, create a process to gracefully shutdown from cmd
	conn.Close()
	listener.Close()
	log.Println("Shut Down Complete...")
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
