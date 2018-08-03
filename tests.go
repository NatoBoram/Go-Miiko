package main

import (
	"database/sql"
	"fmt"
	"os"
)

func testInsertWelcomeChannel() {

	// Get guild
	guild, err := session.Guild("365011717375262720")
	if err != nil {
		fmt.Println("Couldn't get a guild.")
		fmt.Println(err.Error())
		return
	}

	// Get channel
	channel, err := session.Channel("365011718025248769")
	if err != nil {
		fmt.Println("Couldn't get a channel.")
		fmt.Println(err.Error())
		return
	}

	// Insert Welcome Channel
	res, err := insertWelcomeChannel(guild, channel)
	if err != nil {
		fmt.Println("Couldn't insert a welcome channel.")
		fmt.Println(err.Error())
		return
	}

	// Get last insert ID
	liid, err := res.LastInsertId()
	if err != nil {
		fmt.Println("Couldn't get the last insert ID.")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("LastInsertId :", liid)

	// Get rows affected
	ra, err := res.RowsAffected()
	if err != nil {
		fmt.Println("Couldn't get the rows affected.")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("RowsAffected :", ra)

	os.Exit(0)
}

func testSelectNothing() {

	// Get guild
	guild, err := session.Guild("365011717375262720")
	if err != nil {
		fmt.Println("Couldn't get a guild.")
		fmt.Println(err.Error())
		return
	}

	// Select a presentation channel
	pc, err := selectPresentationChannel(guild)
	if err != nil {
		fmt.Println("Couldn't select this presentation channel.")
		fmt.Println(err.Error())

		if err == sql.ErrNoRows {
			fmt.Println("The error is", "sql: no rows in result set")
		}

		return
	}

	fmt.Println("Presentation Channel :", pc)
	os.Exit(0)
}
