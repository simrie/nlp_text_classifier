package db_mongo

import (
	"context"
	"net/http"
	"nlp_text_classifier/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPeople() ([]types.Person, int, error) {
	var people []types.Person

	var client *mongo.Client

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)

	collection := client.Database("thepolyglotdeveloper").Collection("people")

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
