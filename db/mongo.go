package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
)

type Room struct {
	sess         *mgo.Session
	fabrcc       *mgo.Database
	Users        *mgo.Collection
	Transactions *mgo.Collection
}

// Returns a new room
func InitRoom(uri string) Room {
	sess, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error: %v\n", err)
		os.Exit(1)
	}

	// Close session in case mongo fails
	defer func() {
		if r := recover(); r != nil {
			sess.Close()
		}
	}()

	sess.SetSafe(&mgo.Safe{})
	fabrcc := sess.DB("iou-fabrcc")
	users := fabrcc.C("users")
	transactions := fabrcc.C("transactions")

	fmt.Println("Connected to mongohq")

	return Room{sess, fabrcc, users, transactions}
}

// Closes the room
func CloseRoom(room Room) {
	room.sess.Close()
}
