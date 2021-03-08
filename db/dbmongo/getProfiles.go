package dbmongo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"nlp_text_classifier/profile"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
GetProfiles returns a set of profiles from Mongo
*/
func (p Pool) GetProfiles(parentCtx context.Context, dbName string) ([]profile.Profile, int, error) {
	var profiles []profile.Profile

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovering from panic in GetProfiles error is: %v \n", r)
		}
	}()

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

	collection := client.Database(dbName).Collection("profiles")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var profile profile.Profile
		cursor.Decode(&profile)
		profiles = append(profiles, profile)
	}
	if err := cursor.Err(); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return profiles, http.StatusOK, nil
}
