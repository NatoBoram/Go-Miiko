package main

import (
	"database/sql"
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
	insert, err := tx.Prepare("insert into `pins`(`server`, `channel`, `member`, `message`) values(?, ?, ?, ?)")
	if err != nil {
		fmt.Println("Couldn't prepare pins.")
		fmt.Println(err.Error())
		return
	}
	defer insert.Close()

	// For each guilds
	for _, guild := range s.State.Guilds {

		// Status
		statusGuild, err := setManualStatus(s, "vérifier "+guild.Name)
		if err != nil {
			fmt.Println("Couldn't set the status manually to checking guild.")
			fmt.Println(err.Error())
			continue
		}

		// For each channels
		for _, channel := range guild.Channels {

			// Ignore non-text channels
			if channel.Type != discordgo.ChannelTypeGuildText {
				continue
			}

			// Pins
			pins, err := s.ChannelMessagesPinned(channel.ID)
			if err != nil {
				// Probably missing permission. Logging it isn't useful.
				// printDiscordError("Couldn't get a channel's pins.", guild, channel, nil, nil, err)
				continue
			}

			// Status
			statusChannel, err := setManualStatus(s, "vérifier #"+channel.Name)
			if err != nil {
				fmt.Println("Couldn't set the status manually to checking channel.")
				fmt.Println(err.Error())
				continue
			}

			// For each pin
			for _, message := range pins {

				// Insert it!
				_, err := insert.Exec(guild.ID, channel.ID, message.Author.ID, message.ID)
				if err != nil {
					fmt.Println("Couldn't execute a pin.")
					fmt.Println(err.Error())
				}

				// Check if it wasn't already sent to the hall of fame
				_, err = selectMessagesFamed(message)
				if err == sql.ErrNoRows {

					// Fame it!
					// go savePin(s, guild, message)

				} else if err != nil {
					fmt.Println("Couldn't select a message inside the hall of fame.")
					fmt.Println(err.Error())
				}
			}

			// Remove status for this channel
			_, err = deleteStatus(statusChannel)
			if err != nil {
				fmt.Println("Couldn't delete a manual status.")
				fmt.Println(err.Error())
			}
			go refreshStatus(s)
		}

		// Remove status for this guild
		_, err = deleteStatus(statusGuild)
		if err != nil {
			fmt.Println("Couldn't delete a manual status.")
			fmt.Println(err.Error())
		}
		go refreshStatus(s)
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

	// Status!
	err = setStatus(s, "démarrer une belle journée")
	// There's some Twitch.tv hard-coding in Discord. Shame!
	// If the hard-coding ever fades away, then I'll probably include this here.
	// err = s.UpdateStreamingStatus(0, "GitLab", "https://gitlab.com/NatoBoram/Go-Miiko/")
	if err != nil {
		fmt.Println("Couldn't clear the status to beginning a beautiful day.")
		fmt.Println(err.Error())
	}

	// Now that pins are refreshed, it's time to refresh the hall of fame.
	go refreshHallOfFame(s)
}

func refreshHallOfFame(s *discordgo.Session) {

	// For each guilds
	for _, guild := range s.State.Guilds {
		go refreshHallOfFameGuild(s, guild)
	}
}

func refreshHallOfFameGuild(s *discordgo.Session, g *discordgo.Guild) {

	// Select all pins
	rows, err := selectPins(g)
	if err != nil {
		fmt.Println("Couldn't fetch pins for an entire server.")
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	// For each pins
	for rows.Next() {

		var (
			channelID string
			messageID string
		)

		err := rows.Scan(&channelID, &messageID)
		if err != nil {
			fmt.Println("Couldn't fetch the next guild's pins.")
			fmt.Println(err.Error())
			continue
		}

		// Actual process
		message, err := s.ChannelMessage(channelID, messageID)
		if err != nil {
			printDiscordError("Couldn't get a message from a pin.", g, nil, nil, nil, err)
			continue
		}

		// Check before faming a message
		_, err = selectMessagesFamed(message)
		if err == sql.ErrNoRows {

			// Fame it!
			go savePin(s, g, message)

		} else if err != nil {
			fmt.Println("Couldn't select a famed message.")
			fmt.Println(err.Error())
			continue
		}
	}

	// Check for errors
	err = rows.Err()
	if err != nil {
		fmt.Println("Something went wrong when refreshing the hall of fame.")
		fmt.Println(err.Error())
	}
}

func createTables() (res sql.Result, err error) {

	// Declare tables to create
	functs := [...]func() (res sql.Result, err error){
		createTableChannels,
		createTableRoles,
		createTableSAR,
		createTablePin,
		createTableMinimumReactions,
		createTableStatus,
		createViewPinsCount,
		createTableMessageFame,
	}

	// Create the tables
	for _, funct := range functs {
		res, err = funct()
		if err != nil {
			return
		}
	}

	return
}
