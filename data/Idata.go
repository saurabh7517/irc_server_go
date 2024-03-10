package data

import obj "irc_server/objects"

type DataAccessInterface interface {
	getOneUser(username string, password string) (*obj.User, error)
	getUserList() []obj.User
	isUserPresent(username string) bool
	insertUser(username string, password string) bool
}
