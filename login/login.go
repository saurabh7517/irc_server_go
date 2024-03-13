package login

import (
	"errors"
	data "irc_server/data"
	obj "irc_server/objects"
	util "irc_server/pkg"
	"log"

	"github.com/google/uuid"
)

func HandleReqRes(incomingMsg *obj.Message) []byte {
	var respnoseBytes []byte
	var response *obj.Response = &obj.Response{Status: "SUCCESS", Msg: "UNKNOWN", Token: ""}
	userDto, err := extractDto(incomingMsg)
	if err != nil {
		log.Println(err)
		response = &obj.Response{Status: "SUCCESS", Msg: err.Error(), Token: ""}
	} else {
		//Get Database connection
		var dataConn data.DataAccessInterface = data.GetDataSource()
		_, err := dataConn.GetOneUser(userDto.Username, userDto.Password)
		if err != nil {
			log.Println("User not found !")
			response = &obj.Response{Status: "SUCCESS", Msg: err.Error(), Token: ""}
		} else {
			//add user to cache
			//userDto is a vaid user
			//Generate a token for the logged in user
			token := processLogin(*userDto)
			response = &obj.Response{Status: "SUCCESS", Msg: "SUCCESS", Token: token}
		}
	}
	respnoseBytes = util.EncodeResponse(response)
	return respnoseBytes
}

func processLogin(userDto obj.UserDto) string {
	uuid := uuid.NewString()

	var userSession obj.UserSession = obj.UserSession{Username: userDto.Username,
		Token: uuid, HostIp: userDto.HostIp, HostPort: userDto.HostPort}
	util.AddUserToCache(userSession)
	return uuid
}

func extractUserDto(incomingMsg *obj.Message) (*obj.User, error) {
	if len(incomingMsg.User.Username) > 0 && len(incomingMsg.User.Password) > 0 {
		return incomingMsg.User, nil
	} else {
		return nil, errors.New("Invalid username")
	}
}

func extractHostAddressDto(incomingMsg *obj.Message) (*obj.HostAddress, error) {
	if len(incomingMsg.HostAddress.HostIp) > 0 && len(incomingMsg.HostAddress.HostPort) > 0 {
		return incomingMsg.HostAddress, nil
	} else {
		return nil, errors.New("Invalid HostAddress")
	}
}

func extractDto(incomingMsg *obj.Message) (*obj.UserDto, error) {
	validUser, err := extractUserDto(incomingMsg)
	if err != nil {
		return nil, err
	} else {
		validHostAddress, err := extractHostAddressDto(incomingMsg)
		if err != nil {
			return nil, err
		}
		return &obj.UserDto{Username: validUser.Username, Password: validUser.Password,
			HostIp:   validHostAddress.HostIp,
			HostPort: validHostAddress.HostPort}, nil
	}
}
