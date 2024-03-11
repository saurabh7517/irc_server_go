package data

import (
	"errors"
	obj "irc_server/objects"
)

type inMemDb struct {
	userList []obj.User
}

var inMemDbList *inMemDb = &inMemDb{userList: nil}

func initializeInMemDb() {
	//create users
	var userList []obj.User
	var user1 obj.User = obj.User{Username: "kumar", Password: "pass"}
	var user2 obj.User = obj.User{Username: "foo", Password: "bar"}

	userList = append(userList, user1)
	userList = append(userList, user2)
	inMemDbList.userList = userList

}

func getInMemDBConnection() *inMemDb {
	if len(inMemDbList.userList) > 0 {
		// do nothing
		return inMemDbList
	} else {
		initializeInMemDb()
		return inMemDbList
	}
}

// getOneUser(username string, password string) obj.User
// getUserList() []obj.User
// isUserPresent(username string) bool

func (inMemData *inMemDb) GetOneUser(username string, password string) (*obj.User, error) {
	var foundUser *obj.User = &obj.User{}
	for _, user := range inMemData.userList {
		if user.Username == username && user.Password == password {
			foundUser.Username = username
			foundUser.Password = password
			return foundUser, nil
		}
	}
	return nil, errors.New("user not found")
}

func (inMemData *inMemDb) GetUserList() []obj.User {
	return inMemData.userList
}

func (inMemData *inMemDb) IsUserPresent(username string) bool {
	for _, user := range inMemData.userList {
		if user.Username == username {
			return true
		}
	}
	return false
}

func (inMemData *inMemDb) InsertUser(username string, password string) bool {
	inMemData.userList = append(inMemData.userList, obj.User{Username: username, Password: password})
	return true
}
