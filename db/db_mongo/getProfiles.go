package db_mongo

import (
	"context"
	"errors"
	"net/http"
	"nlp_text_classifier/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (p Pool) GetProfiles(parentCtx context.Context, dbName string) ([]types.Person, int, error) {
	var people []types.Person

	c, err := p.Borrow()
	if err != nil {
		return nil, 0, err
	}
	// assert client as *mongo.Client
	client, ok := c.(*mongo.Client)
	if !ok {
		return nil, 0, errors.New("requires *mongo.Client")
	}
	defer p.Restock(client)

	ctx, cancel := context.WithTimeout(parentCtx, 10*time.Second)
	defer cancel()

	collection := client.Database(dbName).Collection("people")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var person types.Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return people, http.StatusOK, nil
}
