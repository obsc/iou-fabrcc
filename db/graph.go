package db

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Graph struct {
	Id        bson.ObjectId   "_id"
	Nodes     map[string]Node "nodes"
	CreatedAt time.Time       "createdAt"
	UpdatedAt time.Time       "updatedAt"
}

type Node struct {
	Id    bson.ObjectId   "userId"
	Worth int             "worth"
	Edges map[string]Edge "edges"
}

type Edge struct {
	Value int "value"
}

func InitGraph() Graph {
	graph := Graph{
		Id:        bson.NewObjectId(),
		Nodes:     nil,
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
	graph := GetGraph()
	node := Node{
		Id:    user.Id,
		Worth: 0,
		Edges: nil}

	err := room.graph.UpdateId(graph.Id, bson.M{
		"$set": bson.M{"nodes." + user.Id.Hex(): node, "updatedAt": bson.Now()}})

	logError(err)
}

func GraphAddTransaction(transaction Transaction) {
	graph := GetGraph()
	source := graph.Nodes[transaction.SourceId.Hex()]
	sink := graph.Nodes[transaction.SinkId.Hex()]

	sourceEdge := Edge{
		Value: source.Edges[transaction.SinkId.Hex()].Value - transaction.Value}
	sourceWorth := source.Worth - transaction.Value

	sinkEdge := Edge{
		Value: sink.Edges[transaction.SourceId.Hex()].Value + transaction.Value}
	sinkWorth := sink.Worth + transaction.Value

	sourceString := "nodes." + transaction.SourceId.Hex() + "."
	sinkString := "nodes." + transaction.SinkId.Hex() + "."

	err := room.graph.UpdateId(graph.Id, bson.M{
		"$set": bson.M{
			sourceString + "edges." + transaction.SinkId.Hex(): sourceEdge,
			sourceString + "worth":                             sourceWorth,
			sinkString + "edges." + transaction.SourceId.Hex(): sinkEdge,
			sinkString + "worth":                               sinkWorth,
			"updatedAt":                                        bson.Now()}})

	logError(err)
}
