package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
)

type Room struct {
	sess         *mgo.Session
	fabrcc       *mgo.Database
	users        *mgo.Collection
	transactions *mgo.Collection
	graph        *mgo.Collection
}

var room Room

var updateTimeQuery = bson.M{"$set": bson.M{"updatedAt": bson.Now()}}

func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Initializes a new room
func InitRoom(uri string) {
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

	room = Room{
		sess:         sess,
		fabrcc:       fabrcc,
		users:        fabrcc.C("users"),
		transactions: fabrcc.C("transactions"),
		graph:        fabrcc.C("graph")}

	fmt.Println("Connected to mongohq")
}

// Closes the room
func CloseRoom() {
	room.sess.Close()
}
