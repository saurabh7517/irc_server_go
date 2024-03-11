package pkg

import (
	obj "irc_server/objects"
	"log"

	"google.golang.org/protobuf/proto"
)

func EncodeResponse(response *obj.Response) []byte {
	msgBytes, err := proto.Marshal(response)
	if err != nil {
		log.Println("Error in marshaling data")
	}
	return msgBytes
}

func EncodeMessage(message *obj.Message) []byte {
	msgBytes, err := proto.Marshal(message)
	if err != nil {
		log.Println("Error marshalling object")
	}
	return msgBytes
}
