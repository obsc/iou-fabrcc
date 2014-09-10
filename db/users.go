package db

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id        bson.ObjectId        "_id"
	Name      string               "name"
	In        map[string]TransList "in"
	Out       map[string]TransList "out"
	CreatedAt time.Time            "createdAt"
	UpdatedAt time.Time            "updatedAt"
}

type TransList []bson.ObjectId

func GetUsers(query interface{}, limit int) []User {
	results := []User{}
	err := room.users.Find(query).Sort("-createdAt").Limit(limit).All(&results)
	logError(err)

	return results
}

func IterUsers(query interface{}, fn func(User)) error {
	result := User{}
	iter := room.users.Find(query).Sort("-createdAt").Iter()

	for iter.Next(&result) {
		fn(result)
	}

	return iter.Close()
}

func GetUserNameMap(query interface{}) map[bson.ObjectId]string {
	m := make(map[bson.ObjectId]string)
	IterUsers(query, func(user User) {
		m[user.Id] = user.Name
	})
	return m
}

func AddUser(user User) {
	err := room.users.Insert(user)
	GraphAddUser(user)
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
		"$push": bson.M{"out." + t.Hex(): trans}})
	logError(err)
	err = room.users.UpdateId(s, updateTimeQuery)
	logError(err)

	err = room.users.UpdateId(t, bson.M{
		"$push": bson.M{"in." + s.Hex(): trans}})
	logError(err)
	err = room.users.UpdateId(t, updateTimeQuery)
	logError(err)
}
