package main

import (
	"fmt"
	"nlp_text_classifier/db"
	"nlp_text_classifier/db/dbmongo"
	"nlp_text_classifier/server"
)

func main() {
	//initialize a pool of connections
	var connectionString = "mongodb://localhost:27017"
	var pool db.Pool
	var err error
	pool, err = dbmongo.Init(5, connectionString)
	if err != nil {
		fmt.Println(err)
		return
	}

	server.StartRouter(pool)
}
