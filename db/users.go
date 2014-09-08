package db

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id        bson.ObjectId            "_id"
	Name      string                   "name"
	In        map[bson.ObjectId]DiEdge "in"
	Out       map[bson.ObjectId]DiEdge "out"
	CreatedAt time.Time                "createdAt"
	UpdatedAt time.Time                "updatedAt"
}

type DiEdge []bson.ObjectId

func GetUsers(query interface{}, limit int) []User {
	results := []User{}
	err := room.users.Find(query).Limit(limit).All(&results)
	handleError(err)

	return results
}

func IterUsers(query interface{}, fn func(User)) error {
	result := User{}
	iter := room.users.Find(query).Iter()

	for iter.Next(&result) {
		fn(result)
	}

	return iter.Close()
}

func AddUser(user User) {
	err := room.users.Insert(user)
	handleError(err)
}

func AddUserByName(name string) {
	AddUser(User{
		Id:        bson.NewObjectId(),
		Name:      name,
		In:        nil,
		Out:       nil,
		CreatedAt: bson.Now(),
		UpdatedAt: bson.Now()})
}

func UpdateUserTransaction(trans bson.ObjectId, s bson.ObjectId, t bson.ObjectId) {
	err := room.users.UpdateId(s, bson.M{
		"$push": bson.M{"out." + t.String(): trans}})
	handleError(err)
	err = room.users.UpdateId(s, updateTimeQuery)
	handleError(err)

	err = room.users.UpdateId(t, bson.M{
		"$push": bson.M{"in." + s.String(): trans}})
	handleError(err)
	err = room.users.UpdateId(t, updateTimeQuery)
	handleError(err)
}
