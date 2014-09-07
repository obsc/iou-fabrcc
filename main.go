package main

import (
	"fmt"
	"github.com/obsc/iou-fabrcc/db"
	"net/http"
	"os"
)

const DEFAULT_URI string = "mongodb://fabrcc:fabrcc#420rekt@kahana.mongohq.com:10097/iou-fabrcc"
const DEFAULT_PORT string = "5000"

var URI string = os.Getenv("MONGOHQ_URL")
var PORT string = os.Getenv("PORT")
var ROOM db.Room

func main() {
	routes()

	if URI == "" {
		ROOM = db.InitRoom(DEFAULT_URI)
	} else {
		ROOM = db.InitRoom(URI)
	}

	listen(PORT)
}

func routes() {
	http.HandleFunc("/", root)
}

func listen(port string) {
	if port == "" {
		port = DEFAULT_PORT
	}

	fmt.Println("listening...")
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		panic(err)
	}
}

func root(res http.ResponseWriter, req *http.Request) {
	count, _ := ROOM.Users.Count()
	fmt.Fprintln(res, "hello world!", count)
	// fmt.Fprintln(res, "hello world!")
}
