package common

var UserMap map[string]bool = make(map[string]bool)

func AddLogedInUser(nickname string) {
	UserMap[nickname] = true
}

func LogoutUser(nickName string) {
	delete(UserMap, nickName)

}

func IsUserOnServer(nickName string) bool {
	_, exists := UserMap[nickName]
	return exists
}
