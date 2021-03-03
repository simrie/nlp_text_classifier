package main

import (
	"fmt"
	"nlp_text_classifier/db/db_mongo"
	"nlp_text_classifier/server"
)

func main() {
	//initialize a pool of connections
	var connectionString = "mongodb://localhost:27017"
	pool, err := db_mongo.Init(5, connectionString)
	if err != nil {
		fmt.Println(err)
		return
	}

	server.StartRouter(pool)
}
