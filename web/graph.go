package web

import (
	"fmt"
	"github.com/obsc/iou-fabrcc/db"
	"net/http"
)

func graph(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Nothing to see here folks")
}

func graphJson(w http.ResponseWriter, r *http.Request) {
	type nodeJsonable struct {
		Worth int
		Edges map[string]int
	}
	graphJsons := make(map[string]nodeJsonable)

	graph := db.GetGraph()
	for id, node := range graph.Nodes {
		edges := make(map[string]int)
		for id, value := range node.Edges {
			edges[id] = value.Value
		}

		graphJsons[id] = nodeJsonable{
			Worth: node.Worth,
			Edges: edges}
	}

	printJson(w, graphJsons)
}
