package main

import (
	"strconv"
)

func toConnectionString(database Database) string {
	return database.User + ":" + database.Password + "@tcp(" + database.Address + ":" + strconv.Itoa(database.Port) + ")/" + database.Database
}
