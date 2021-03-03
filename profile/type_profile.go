package profile

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
WordsSeen shows original words that were stemmed and their counts
*/
type WordSeen struct {
	Word string `json:"word"`
	Seen int    `json:"seen"`
}

/*
Block has the minified stem of words extracted
*/
type Block struct {
	MiniStem string      `json:"mini_stem"`
	Source   []WordSeen `json:"source"`
	Weight   int         `json:"weight"`
	Count    int         `json:"count"`
}

/*
BlockMap is for organizing Blocks
*/
type BlockMap = map[string]Block

/*
Profile is extracted from RawDoc Text
*/
type Profile struct {
	Name   string  `json:"name" bson:"name,omitempty"`
	Tag    string  `json:"tag" bson:"tag,omitempty"`
	Blocks []Block `json:"blocks,omitempty" bson:"blocks,omitempty"`
}

/*
ProfileMongo is Profile but for Mongo
*/
type ProfileMongo struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Profile
}
