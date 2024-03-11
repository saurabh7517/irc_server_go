package register

import (
	"errors"
	data "irc_server/data"
	obj "irc_server/objects"
	util "irc_server/pkg"
	"log"
)

func HandleReqRes(incomingMsg *obj.Message) []byte {
	var respnoseBytes []byte
	var response *obj.Response = &obj.Response{Status: "SUCCESS", Msg: "UNKNOWN", Token: ""}
	user, err := extractUserDto(incomingMsg)
	if err != nil {
		log.Println(err)
	} else {
		//Get Database connection
		var dataConn data.DataAccessInterface = data.GetDataSource()

		if dataConn.IsUserPresent(user.Username) { //check if there is an existing username in database
			response = &obj.Response{Status: "SUCCESS", Msg: "ALREADY_EXISTS", Token: ""}
			respnoseBytes = util.EncodeResponse(response)

		} else {
			//insert a new row in the database
			processRegistration(user.GetUsername(), user.GetPassword())
			response = &obj.Response{Status: "SUCCESS", Msg: "CREATED", Token: ""}
			respnoseBytes = util.EncodeResponse(response)
		}
	}
	return respnoseBytes
}

func extractUserDto(incomingMsg *obj.Message) (*obj.User, error) {
	if len(incomingMsg.User.Username) > 0 && len(incomingMsg.User.Password) > 0 {
		return incomingMsg.User, nil
	} else {
		return nil, errors.New("Invalid username")
	}
}

func processRegistration(username string, password string) bool {
	var dataConn data.DataAccessInterface = data.GetDataSource()
	return dataConn.InsertUser(username, password)
}
