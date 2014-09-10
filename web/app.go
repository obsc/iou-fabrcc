package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/obsc/iou-fabrcc/db"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

type WebApp struct {
	Router *mux.Router
}

func newWebApp() *WebApp {
	app := &WebApp{
		Router: mux.NewRouter(),
	}
	return app
}

var App *WebApp = newWebApp()

func (app *WebApp) SetRoutes() {
	// request multiplexer
	app.Router.HandleFunc("/users/new", newUser).Name("newUser").Methods("POST")
	app.Router.HandleFunc("/users/json", usersJson).Name("usersJson").Methods("GET")
	app.Router.HandleFunc("/users", users).Name("users").Methods("GET")

	app.Router.HandleFunc("/transactions/new", newTransaction).Name("newTransaction").Methods("POST")
	app.Router.HandleFunc("/transactions/json", transactionsJson).Name("transactionsJson").Methods("GET")
	app.Router.HandleFunc("/transactions", transactions).Name("transactions").Methods("GET")

	app.Router.HandleFunc("/graph/json", graphJson).Name("graphJson").Methods("GET")
	app.Router.HandleFunc("/graph", graph).Name("graph").Methods("GET")

	app.Router.HandleFunc("/", index).Name("index")
}

func (app *WebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.Router.ServeHTTP(w, r)
}

func moneyFilter(i int) string {
	return strconv.Itoa(i)
}

func printJson(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	if err == nil {
		fmt.Fprintf(w, "%s", b)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

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

func transactions(w http.ResponseWriter, r *http.Request) {
	userNameMap := db.GetUserNameMap(nil)

	db.IterTransactions(nil, func(trans db.Transaction) {
		fmt.Fprintf(w, "%s owes %s %s because of: %s\n",
			userNameMap[trans.SourceId], userNameMap[trans.SinkId],
			moneyFilter(trans.Value), trans.Reason)
	})
}

func transactionsJson(w http.ResponseWriter, r *http.Request) {
	type transJsonable struct {
		TransactionId string
		SourceId      string
		SinkId        string
		Value         int
		Reason        string
	}
	transJsons := make([]transJsonable, 0)

	db.IterTransactions(nil, func(trans db.Transaction) {
		t := transJsonable{
			TransactionId: trans.Id.Hex(),
			SourceId:      trans.SourceId.Hex(),
			SinkId:        trans.SinkId.Hex(),
			Value:         trans.Value,
			Reason:        trans.Reason}
		transJsons = append(transJsons, t)
	})

	printJson(w, transJsons)
}

func newTransaction(w http.ResponseWriter, r *http.Request) {
	source := r.FormValue("source")
	sink := r.FormValue("sink")
	valuestr := r.FormValue("value")
	reason := r.FormValue("reason")

	if source != sink && bson.IsObjectIdHex(source) && bson.IsObjectIdHex(sink) {
		value, err := strconv.Atoi(valuestr)
		if value < 0 {
			value = -value
			source, sink = sink, source
		}
		if err == nil && reason != "" {
			db.AddTransactionByData(bson.ObjectIdHex(source), bson.ObjectIdHex(sink), value, reason)
		}
	}
}

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
