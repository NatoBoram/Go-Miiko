package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func refresh(s *discordgo.Session) {

	// Wait for a minute
	time.Sleep(time.Minute)

	// Start
	start := time.Now()
	fmt.Println("Begin refresh :", start.String())

	// Begin
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Couldn't begin a pin transaction.")
		fmt.Println(err.Error())
		return
	}

	// Delete
	result, err := tx.Exec("delete from `pins`;")
	if err != nil {
		fmt.Println("Couldn't delete all pins.")
		fmt.Println(err.Error())
		return
	}

	// Rows Affected
	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Couldn't get all affected pins.")
		fmt.Println(err.Error())
		return
	}

	// Log
	fmt.Println("Deleted", rows, "pins.")

	// Prepare
	insert, err := tx.Prepare("insert into `pins`(`server`, `member`, `message`) values(?, ?, ?)")
	if err != nil {
		fmt.Println("Couldn't prepare pins.")
		fmt.Println(err.Error())
		return
	}
	defer insert.Close()

	// For each guilds
	for _, guild := range s.State.Guilds {
		s.UpdateStatus(0, "vérifier "+guild.Name)

		// For each channels
		for _, channel := range guild.Channels {

			// Ignore non-text channels
			if channel.Type != discordgo.ChannelTypeGuildText {
				continue
			}

			// Pins
			pins, err := s.ChannelMessagesPinned(channel.ID)
			if err != nil {
				fmt.Println("Couldn't get a channel's pins.")
				fmt.Println("Guild :", guild.Name)
				fmt.Println("Channel :", channel.Name)
				fmt.Println(err.Error())
				continue
			}

			s.UpdateStatus(0, "vérifier "+guild.Name+"/"+channel.Name)

			// For each pin
			for _, message := range pins {

				// Insert it!
				_, err := insert.Exec(guild.ID, message.Author.ID, message.ID)
				if err != nil {
					fmt.Println("Couldn't execute a pin.")
					fmt.Println(err.Error())
				}
			}
		}
	}

	// Commit
	err = tx.Commit()
	if err != nil {
		fmt.Println("Couldn't commit a pin transaction.")
		fmt.Println(err.Error())
	}

	// End
	duration := time.Since(start)
	fmt.Println("Finished refresh :", duration.String())
}
