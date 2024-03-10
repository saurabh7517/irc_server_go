package data

import obj "irc_server/objects"

type DataAccessInterface interface {
	GetOneUser(username string, password string) (*obj.User, error)
	GetUserList() []obj.User
	IsUserPresent(username string) bool
	InsertUser(username string, password string) bool
}
