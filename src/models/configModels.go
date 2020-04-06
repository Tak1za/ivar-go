package models

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

type ServerConfiguration struct {
	URL      string
	Username string
	Password string
}

type DatabaseConfiguration struct {
	Index string
}
