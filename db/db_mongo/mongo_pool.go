package db_mongo

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type Pool should implement interface DB_Pool

type Pool struct {
	Connections []*mongo.Client
	Mutex       *sync.Mutex
}

// Simple pooling code was partly inspired by
// https://www.reddit.com/r/golang/comments/63w8u1/restful_api_without_openingclosing_database_for/

func Init(poolSize int, uri string) (Pool, error) {
	if poolSize <= 0 {
		return Pool{}, errors.New("invalid poolSize")
	}
	connections := make([]*mongo.Client, poolSize)
	for count := 0; count < poolSize; count++ {
		fmt.Println("My counter is at", count)
		var client *mongo.Client

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		clientOptions := options.Client().ApplyURI(uri)
		client, _ = mongo.Connect(ctx, clientOptions)
		connections[count] = client
	}
	var pool = Pool{
		Connections: connections,
		Mutex:       &sync.Mutex{},
	}
	return pool, nil
}

func (p Pool) Borrow() (interface{}, error) {
	if len(p.Connections) == 0 {
		return nil, errors.New("Cannot return connection")
	}

	var client *mongo.Client

	p.Mutex.Lock()
	client = p.Connections[0]
	p.Connections = p.Connections[1:]
	p.Mutex.Unlock()
	fmt.Println("borrowed")
	return client, nil
}

func (p Pool) Restock(c interface{}) error {
	// assert client as *mongo.Client
	client, ok := c.(*mongo.Client)
	if !ok {
		return errors.New("requires *mongo.Client")
	}
	p.Mutex.Lock()
	p.Connections = append(p.Connections, client)
	p.Mutex.Unlock()
	fmt.Println("restocked")
	return nil
}
