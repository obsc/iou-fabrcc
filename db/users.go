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
	room.users.Find(query).Limit(limit).All(&results)
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
	room.users.Insert(user)
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
	room.users.UpdateId(s, bson.M{
		"$push": bson.M{"out." + t.String(): trans}})
	room.users.UpdateId(s, updateTimeQuery)

	room.users.UpdateId(t, bson.M{
		"$push": bson.M{"in." + s.String(): trans}})
	room.users.UpdateId(t, updateTimeQuery)
}
