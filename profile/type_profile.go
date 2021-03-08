package profile

/*
Profile is extracted from RawDoc Text
bson values are specifical to MongoDB
and should be ignored if dealing with another DB
*/
type Profile struct {
	Name   string  `json:"name" bson:"name,omitempty"`
	Tag    string  `json:"tag" bson:"tag,omitempty"`
	DbKey  string  `json:"db_key,omitempty" bson:"_id,omitempty"`
	Blocks []Block `json:"blocks,omitempty" bson:"blocks,omitempty"`
}
