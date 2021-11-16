package database

type ConfigDB struct {
	Servers []string
	User string
	Password string
	Port int
	Database string
	MaxIdle int
	MaxOpen int
	UseSSL bool
	ShowSQL bool
}
