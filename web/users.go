package web

import (
	"fmt"
	"github.com/obsc/iou-fabrcc/db"
	"net/http"
)

func users(w http.ResponseWriter, r *http.Request) {
	graph := db.GetGraph()
	userNameMap := db.GetUserNameMap(nil)

	for id, node := range graph.Nodes {
		fmt.Fprintf(w, "%s          %s: %s\n",
			id, userNameMap[node.Id], moneyFilter(node.Worth))
	}
}

func usersJson(w http.ResponseWriter, r *http.Request) {
	type usersJsonable struct {
		UserId string
		Name   string
	}
	usersJsons := make([]usersJsonable, 0)

	db.IterUsers(nil, func(user db.User) {
		u := usersJsonable{
			UserId: user.Id.Hex(),
			Name:   user.Name}
		usersJsons = append(usersJsons, u)
	})

	printJson(w, usersJsons)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	var name string = r.FormValue("name")
	if name != "" {
		db.AddUserByName(name)
	}
}
