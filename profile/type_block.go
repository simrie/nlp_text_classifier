package profile

/*
WordSeen shows original words that were stemmed and their counts
*/
type WordSeen struct {
	Word string `json:"word"`
	Seen int    `json:"seen"`
}

/*
Block has the minified stem of words extracted
*/
type Block struct {
	MiniStem string     `json:"mini_stem"`
	Source   []WordSeen `json:"source"`
	Weight   int        `json:"weight"`
	Count    int        `json:"count"`
}

/*
BlockMapType is for Mapped Blocks
*/
type BlockMapType struct {
	BlockMap map[string]Block `json:"block_map,omitempty"`
}
