package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

const STATIC_DIR = "web/static"
const TEMPLATE_DIR = "web/templates"

var TEMPLATE_LIST = []string{"index.html"}

type WebApp struct {
	Router      *mux.Router
	StaticDir   string
	TemplateDir string
	Templates   map[string]*template.Template
}

func parseTemplates() map[string]*template.Template {
	templates := make(map[string]*template.Template)

	for _, filename := range TEMPLATE_LIST {
		templates[filename] = template.Must(
			template.ParseFiles(filepath.Join(TEMPLATE_DIR, filename)))
	}

	return templates
}

func newWebApp() *WebApp {
	app := &WebApp{
		Router:      mux.NewRouter(),
		StaticDir:   STATIC_DIR,
		TemplateDir: TEMPLATE_DIR,
		Templates:   parseTemplates()}
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
	tmpl := App.Templates["index.html"]

	tmpl.Execute(w, nil)
}
