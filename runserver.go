package main

import (
	"github.com/obsc/iou-fabrcc/db"
	"github.com/obsc/iou-fabrcc/web"
	"log"
	"net/http"
	"os"
)

const (
	DEFAULT_URI  string = "mongodb://fabrcc:fabrcc#420rekt@kahana.mongohq.com:10097/iou-fabrcc"
	DEFAULT_PORT string = "5000"
)

var PORT string

func main() {
	initVars()
	web.App.SetRoutes()
	http.Handle("/", web.App)

	listen(PORT)
}

func initVars() {
	if os.Getenv("PORT") == "" {
		log.Println("PORT environment variable not set, defaulting to ", DEFAULT_PORT)
		PORT = ":" + DEFAULT_PORT
	} else {
		PORT = ":" + os.Getenv("PORT")
	}

	if os.Getenv("MONGOHQ_URL") == "" {
		log.Println("MONGOHQ_URL environment variable not set")
		db.InitRoom(DEFAULT_URI)
	} else {
		db.InitRoom(os.Getenv("MONGOHQ_URL"))
	}
}

func listen(port string) {
	log.Println("Starting server on: ", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal("An error occured when trying to start server: \n", err)
	}
}
