package web

import (
	"fmt"
	"github.com/obsc/iou-fabrcc/db"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

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
