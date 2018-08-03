package main

func (discord Discord) isEmpty() bool {
	return discord.MasterID == "" || discord.Token == ""
}

func (database Database) isEmpty() bool {
	return database.Address == "" ||
		database.Database == "" ||
		database.Password == "" ||
		database.Port == 0 ||
		database.User == ""
}
