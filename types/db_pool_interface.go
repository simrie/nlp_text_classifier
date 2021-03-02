package types

import (
	"context"
)

type DB_Pool interface {
	Borrow() (interface{}, error)
	Restock(interface{}) error
	GetPeople(parentCtx context.Context, dbName string) ([]Person, int, error)
}
