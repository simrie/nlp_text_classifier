package db

import (
	"context"
	"nlp_text_classifier/profile"
)

/*
Pool defines database methods partly inspired by
   https://www.reddit.com/r/golang/comments/63w8u1/restful_api_without_openingclosing_database_for/

By using an interface to define database methods
the APIs could be switched to a different database
by passing a different DB_Pool inmplementation

*/
type Pool interface {
	Borrow() (interface{}, error)
	Restock(interface{}) error
	GetDatabases(parentCtx context.Context) ([]string, int, error)
	GetProfiles(parentCtx context.Context, dbName string) ([]profile.Profile, int, error)
	GetProfile(parentCtx context.Context, dbName string, idStr string) (profile.Profile, int, error)
	StoreProfiles(parentCtx context.Context, dbName string, profiles []interface{}) (int, int, error)
}
