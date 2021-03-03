package types

import (
	"context"
)

// By using an interface to define database methods
// the APIs could be switched to a different database
// by passing a different DB_Pool inmplementation

type DB_Pool interface {
	Borrow() (interface{}, error)
	Restock(interface{}) error
	GetDatabases(parentCtx context.Context) ([]string, int, error)
	GetProfiles(parentCtx context.Context, dbName string) ([]Person, int, error)
	GetProfile(parentCtx context.Context, dbName string, idStr string) (Person, int, error)
	StoreProfiles(parentCtx context.Context, dbName string, people []interface{}) (int, int, error)
}
