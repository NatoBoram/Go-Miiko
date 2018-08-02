package main

import (
	"fmt"
	"os"
)

func testInsertWelcomeChannel() {

	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Couldn't begin a transaction.")
		fmt.Println(err.Error())
		return
	}

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
	res, err := insertWelcomeChannel(tx, guild, channel)
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

	// Rollback
	err = tx.Rollback()
	if err != nil {
		fmt.Println("Couldn't rollback.")
		fmt.Println(err.Error())
		return
	}

	os.Exit(0)
}
