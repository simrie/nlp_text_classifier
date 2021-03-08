package dbmongo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
StoreProfiles implements interface Pool for storing Profiles
*/
func (p Pool) StoreProfiles(parentCtx context.Context, dbName string, profiles []interface{}) (int, int, error) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovering from panic in StoreProfiles error is: %v \n", r)
		}
	}()

	c, err := p.Borrow()
	if err != nil {
		return 0, 0, err
	}
	// assert client as *mongo.Client
	client, ok := c.(*mongo.Client)
	if !ok {
		return 0, 0, errors.New("requires *mongo.Client")
	}
	defer p.Restock(client)

	ctx, cancel := context.WithTimeout(parentCtx, 10*time.Second)
	defer cancel()

	collection := client.Database(dbName).Collection("profiles")

	opts := options.InsertMany().SetOrdered(true)
	result, err := collection.InsertMany(ctx, profiles, opts)
	fmt.Printf("\nresult of insert many %v\n", result)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	insertCount := len(result.InsertedIDs)
	return insertCount, http.StatusOK, nil
}
