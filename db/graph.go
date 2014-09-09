package db

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Graph struct {
	Nodes     map[bson.ObjectId]Node "userId"
	CreatedAt time.Time              "createdAt"
	UpdatedAt time.Time              "updatedAt"
}

type Node struct {
	Id    bson.ObjectId            "userId"
	Worth int                      "worth"
	In    map[bson.ObjectId]DiEdge "in"
	Out   map[bson.ObjectId]DiEdge "out"
}

type DiEdge struct {
	Value int "value"
}

func InitGraph() Graph {
	nodes := make(map[bson.ObjectId]Node)

	graph := Graph{
		Nodes:     nodes,
		CreatedAt: bson.Now(),
		UpdatedAt: bson.Now()}

	err := room.graph.Insert(graph)
	logError(err)

	IterUsers(nil, GraphAddUser)
	IterTransactions(nil, GraphAddTransaction)

	return graph
}

func GetGraph() Graph {
	count, err := room.graph.Count()
	logError(err)

	if count == 0 {
		return InitGraph()
	} else {
		result := Graph{}
		room.graph.Find(nil).One(&result)
		return result
	}
}

func GraphAddUser(user User) {

}

func GraphAddTransaction(transaction Transaction) {

}
