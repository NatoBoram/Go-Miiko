package main

import (
	"database/sql"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func refreshStatus(s *discordgo.Session) {
	status, err := selectStatus()
	if err == sql.ErrNoRows {

		// If there's none, clear it
		err = s.UpdateStatus(0, "")
		if err != nil {
			fmt.Println("Couldn't clear the status.")
			fmt.Println(err.Error())
		}

	} else if err != nil {
		fmt.Println("Couldn't retrieve the latest status.")
		fmt.Println(err.Error())

		// On errors, clear it
		err = s.UpdateStatus(0, "")
		if err != nil {
			fmt.Println("Couldn't clear the status after an error retreiving it.")
			fmt.Println(err.Error())
		}
		return
	}

	// Actual update of the status
	err = s.UpdateStatus(0, status)
	if err != nil {
		fmt.Println("Couldn't update the status")
		fmt.Println(err.Error())
	}
}
