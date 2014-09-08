package db

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id   bson.ObjectId            "_id"
	Name string                   "name"
	In   map[bson.ObjectId]DiEdge "in"
	Out  map[bson.ObjectId]DiEdge "out"
}

type DiEdge []bson.ObjectId

func GetUsers(limit int) []User {
	results := []User{}
	room.users.Find(nil).Limit(limit).All(&results)
	return results
}

func IterUsers(fn func(user User)) error {
	result := User{}
	iter := room.users.Find(nil).Iter()

	for iter.Next(&result) {
		fn(result)
	}

	return iter.Close()
}

func AddUser(user User) {
	room.users.Insert(user)
}

func AddUserByName(name string) {
	AddUser(User{
		Id:   bson.NewObjectId(),
		Name: name,
		In:   nil,
		Out:  nil,
	})
}
