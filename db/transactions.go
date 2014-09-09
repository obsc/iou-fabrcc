package db

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Transaction struct {
	Id        bson.ObjectId "_id"
	SourceId  bson.ObjectId "sourceId"
	SinkId    bson.ObjectId "sinkId"
	Value     int           "value"
	Reason    string        "reason"
	CreatedAt time.Time     "createdAt"
	UpdatedAt time.Time     "updatedAt"
}

func GetTransactions(query interface{}, limit int) []Transaction {
	results := []Transaction{}
	err := room.transactions.Find(query).Limit(limit).All(&results)
	logError(err)

	return results
}

func IterTransactions(query interface{}, fn func(Transaction)) error {
	result := Transaction{}
	iter := room.transactions.Find(query).Iter()

	for iter.Next(&result) {
		fn(result)
	}

	return iter.Close()
}

func AddTransaction(transaction Transaction) {
	err := room.transactions.Insert(transaction)
	logError(err)

	UpdateUserTransaction(transaction.Id, transaction.SourceId, transaction.SinkId)

	GraphAddTransaction(transaction)
}

func AddTransactionByData(s bson.ObjectId, t bson.ObjectId, v int, r string) {
	AddTransaction(Transaction{
		Id:        bson.NewObjectId(),
		SourceId:  s,
		SinkId:    t,
		Value:     v,
		Reason:    r,
		CreatedAt: bson.Now(),
		UpdatedAt: bson.Now()})
}
