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
	app.Router.HandleFunc("/", index).Name("index")
}

func (app *WebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.Router.ServeHTTP(w, r)
}

func index(w http.ResponseWriter, r *http.Request) {
	count, _ := db.ROOM.Users.Count()
	fmt.Fprintln(w, "hello world!", count)
}
