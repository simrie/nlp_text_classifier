package dbmongo

import (
	"context"
	"errors"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
GetDatabases returns a list of databases
*/
func (p Pool) GetDatabases(parentCtx context.Context) ([]string, int, error) {
	var databases []string
	var err error

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

	// TODO:  add filter as 2nd param.  bson.D{} is empty filter.
	databases, err = client.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return databases, http.StatusOK, nil
}
