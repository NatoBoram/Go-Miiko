package main

// Database hosts the bot's database configuration.
type Database struct {
	User     string
	Password string
	Address  string
	Port     int
	Database string
}

// Discord hosts the bot's Discord configuration.
type Discord struct {
	Token    string
	MasterID string
}

// Languages is the structure that holds all the bot's supported languages.
type Languages struct {
	French  string
	English string
}
