package profile

/*
RawDoc holds incoming doc objects
*/
type RawDoc struct {
	Key  string `json:"key"`
	Text string `json:"text"`
	Tag  string `json:"tag"`
}
