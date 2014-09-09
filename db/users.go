package db

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id        bson.ObjectId               "_id"
	Name      string                      "name"
	In        map[bson.ObjectId]TransList "in"
	Out       map[bson.ObjectId]TransList "out"
	CreatedAt time.Time                   "createdAt"
	UpdatedAt time.Time                   "updatedAt"
}

type TransList []bson.ObjectId

func GetUsers(query interface{}, limit int) []User {
	results := []User{}
	err := room.users.Find(query).Limit(limit).All(&results)
	logError(err)

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
	logError(err)
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
	logError(err)
	err = room.users.UpdateId(s, updateTimeQuery)
	logError(err)

	err = room.users.UpdateId(t, bson.M{
		"$push": bson.M{"in." + s.String(): trans}})
	logError(err)
	err = room.users.UpdateId(t, updateTimeQuery)
	logError(err)
}
