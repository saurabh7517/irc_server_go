package pkg

import obj "irc_server/objects"

var UserMap map[string]obj.UserSession = make(map[string]obj.UserSession)

func AddUserToCache(userSession obj.UserSession) {
	UserMap[userSession.Username] = userSession
}

func LogoutUser(username string) {
	delete(UserMap, username)

}

func IsUserOnServer(username string) bool {
	_, exists := UserMap[username]
	return exists
}
