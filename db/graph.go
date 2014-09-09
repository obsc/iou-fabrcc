package db

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Graph struct {
	Id        bson.ObjectId          "_id"
	Nodes     map[bson.ObjectId]Node "nodes"
	CreatedAt time.Time              "createdAt"
	UpdatedAt time.Time              "updatedAt"
}

type Node struct {
	Id    bson.ObjectId          "userId"
	Worth int                    "worth"
	Edges map[bson.ObjectId]Edge "edges"
}

type Edge struct {
	Value int "value"
}

func InitGraph() Graph {
	nodes := make(map[bson.ObjectId]Node)

	graph := Graph{
		Id:        bson.NewObjectId(),
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
	node := Node{
		Id:    user.Id,
		Worth: 0,
		Edges: make(map[bson.ObjectId]Edge)}

	err := room.graph.Update(nil, bson.M{
		"$set": bson.M{"nodes." + user.Id.String(): node, "updatedAt": bson.Now()}})

	logError(err)
}

func GraphAddTransaction(transaction Transaction) {
	graph := GetGraph()
	source := graph.Nodes[transaction.SourceId]
	sink := graph.Nodes[transaction.SinkId]

	sourceEdge := Edge{
		Value: source.Edges[transaction.SinkId].Value - transaction.Value}
	sourceWorth := source.Worth - transaction.Value

	sinkEdge := Edge{
		Value: sink.Edges[transaction.SourceId].Value + transaction.Value}
	sinkWorth := sink.Worth + transaction.Value

	sourceString := "nodes." + transaction.SourceId.String() + "."
	sinkString := "nodes." + transaction.SinkId.String() + "."

	err := room.graph.Update(nil, bson.M{
		"$set": bson.M{
			sourceString + "edges." + transaction.SinkId.String(): sourceEdge,
			sourceString + "worth":                                sourceWorth,
			sinkString + "edges." + transaction.SourceId.String(): sinkEdge,
			sinkString + "worth":                                  sinkWorth,
			"updatedAt":                                           bson.Now()}})

	logError(err)
}
