package objects

type UserSession struct {
	Username string
	Token    string
	HostIp   string
	HostPort string
}

type UserDto struct {
	Username string
	Password string
	HostIp   string
	HostPort string
}
