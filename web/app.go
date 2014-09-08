package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/obsc/iou-fabrcc/db"
	"net/http"
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
	app.Router.HandleFunc("/new/user/{name}", newUser).Name("newUser")
	app.Router.HandleFunc("/", index).Name("index")
}

func (app *WebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.Router.ServeHTTP(w, r)
}

func index(w http.ResponseWriter, r *http.Request) {
	db.IterUsers(nil, func(user db.User) {
		fmt.Fprintln(w, user.Name)
	})
}

func newUser(w http.ResponseWriter, r *http.Request) {
	var name string = mux.Vars(r)["name"]
	db.AddUserByName(name)

	fmt.Fprintln(w, "Added new user: ", name)
}
