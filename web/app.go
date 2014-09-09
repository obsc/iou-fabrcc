package web

import (
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
	app.Router.HandleFunc("/users", users).Name("users").Methods("GET")
	app.Router.HandleFunc("/transactions/new", newTransaction).Name("newTransaction").Methods("POST")
	app.Router.HandleFunc("/transactions", transactions).Name("transactions").Methods("GET")
	app.Router.HandleFunc("/", index).Name("index")
}

func (app *WebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.Router.ServeHTTP(w, r)
}

func moneyFilter(i int) string {
	return strconv.Itoa(i)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func users(w http.ResponseWriter, r *http.Request) {
	db.IterUsers(nil, func(user db.User) {
		fmt.Fprintln(w, user.Name)
	})
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
		fmt.Fprintf(w, "%s owes %s %s because of: %s",
			userNameMap[trans.SourceId], userNameMap[trans.SinkId],
			moneyFilter(trans.Value), trans.Reason)
	})
}

func newTransaction(w http.ResponseWriter, r *http.Request) {
	source := r.FormValue("source")
	sink := r.FormValue("sink")
	valuestr := r.FormValue("value")
	reason := r.FormValue("reason")

	if source != sink && bson.IsObjectIdHex(source) && bson.IsObjectIdHex(sink) {
		value, err := strconv.Atoi(valuestr)
		if err == nil && reason != "" {
			db.AddTransactionByData(bson.ObjectIdHex(source), bson.ObjectIdHex(sink), value, reason)
		}
	}
}
