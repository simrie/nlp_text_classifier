package db_mongo

import (
	"context"
	"net/http"
	"nlp_text_classifier/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetPeople(p *Pool, parentCtx context.Context, dbName string) ([]types.Person, int, error) {
	var people []types.Person

	//var client *mongo.Client
	client, err := p.Borrow()
	if err != nil {
		return nil, 0, err
	}
	defer p.Restock(client)

	ctx, cancel := context.WithTimeout(parentCtx, 10*time.Second)
	defer cancel()

	collection := client.Database(dbName).Collection("people")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return people, http.StatusInternalServerError, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var person types.Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		return people, http.StatusInternalServerError, err
	}
	return people, http.StatusOK, nil
}
