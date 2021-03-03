package db_mongo

import (
	"context"
	"errors"
	"net/http"
	"nlp_text_classifier/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (p Pool) GetProfile(parentCtx context.Context, dbName string, idStr string) (types.Person, int, error) {
	var person types.Person
	var err error
	var docID primitive.ObjectID

	c, err := p.Borrow()
	if err != nil {
		return types.Person{}, 0, err
	}
	// assert client as *mongo.Client
	client, ok := c.(*mongo.Client)
	if !ok {
		return types.Person{}, 0, errors.New("requires *mongo.Client")
	}
	defer p.Restock(client)

	ctx, cancel := context.WithTimeout(parentCtx, 10*time.Second)
	defer cancel()

	collection := client.Database(dbName).Collection("people")

	// useful info on BSON filter and converting string to BSON primitive.ObjectID
	// https://kb.objectrocket.com/mongo-db/how-to-find-a-mongodb-document-by-its-bson-objectid-using-golang-452

	docID, err = primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return types.Person{}, http.StatusInternalServerError, err
	}
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&person)
	if err != nil {
		return types.Person{}, http.StatusInternalServerError, err
	}

	return person, http.StatusOK, nil
}
