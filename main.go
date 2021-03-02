package main

import (
	"fmt"
	"nlp_text_classifier/db_mongo"
	"nlp_text_classifier/server"
)

func main() {
	str := server.Test("nlp_text_classifier starting....")
	fmt.Println(str)

	//initialize a pool of connections
	var connection_string = "mongodb://localhost:27017"
	pool, err := db_mongo.Init(5, connection_string)
	if err != nil {
		fmt.Println(err)
		return
	}

	server.StartRouter(&pool)
}
