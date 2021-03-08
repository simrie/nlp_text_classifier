package dbmongo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"nlp_text_classifier/profile"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
GetProfile retrieves a a profile by the id supplied
*/
func (p Pool) GetProfile(parentCtx context.Context, dbName string, idStr string) (profile.Profile, int, error) {
	var profile profile.Profile = profile.Profile{}
	var err error
	var docID primitive.ObjectID

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovering from panic in GetProfile error is: %v \n", r)
		}
	}()

	c, err := p.Borrow()
	if err != nil {
		return profile, 0, err
	}
	// assert client as *mongo.Client
	client, ok := c.(*mongo.Client)
	if !ok {
		return profile, 0, errors.New("requires *mongo.Client")
	}
	defer p.Restock(client)

	ctx, cancel := context.WithTimeout(parentCtx, 10*time.Second)
	defer cancel()

	collection := client.Database(dbName).Collection("profiles")

	// useful info on BSON filter and converting string to BSON primitive.ObjectID
	// https://kb.objectrocket.com/mongo-db/how-to-find-a-mongodb-document-by-its-bson-objectid-using-golang-452

	docID, err = primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return profile, http.StatusInternalServerError, err
	}
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&profile)
	if err != nil {
		return profile, http.StatusInternalServerError, err
	}
	return profile, http.StatusOK, nil
}
