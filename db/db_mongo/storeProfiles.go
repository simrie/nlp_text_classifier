package db_mongo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (p Pool) StoreProfiles(parentCtx context.Context, dbName string, people []interface{}) (int, int, error) {

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

	collection := client.Database(dbName).Collection("people")

	opts := options.InsertMany().SetOrdered(false)
	result, err := collection.InsertMany(ctx, people, opts)
	fmt.Printf("\nresult of insert many %v\n", result)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	insertCount := len(result.InsertedIDs)

	return insertCount, http.StatusOK, nil
}
