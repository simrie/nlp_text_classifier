package types

import (
	"context"
)

// By using an interface to define database methods
// we could switch between different databases for the backend

type DB_Pool interface {
	Borrow() (interface{}, error)
	Restock(interface{}) error
	GetPeople(parentCtx context.Context, dbName string) ([]Person, int, error)
	GetDatabases(parentCtx context.Context) ([]string, int, error)
}
