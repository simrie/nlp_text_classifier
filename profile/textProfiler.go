package profile

import (
	"nlp_text_classifier/nlp"

	"github.com/jdkato/prose/v3"
)

/*
TextProfiler extracts non-stop word stems from RawDoc.Text
append those as Blocks to a Profile object
*/
func (rawDoc RawDoc) TextProfiler() (Profile, error) {

	var profile Profile
	var proseDoc *prose.Document
	var err error

	normalized, err := nlp.NormalizeText(rawDoc.Text)
	if err != nil {
		return Profile{}, err
	}

	proseDoc, err = nlp.MakeSegmenter(normalized)
	if err != nil {
		return Profile{}, err
	}

	var blockMap = make(map[string]Block)
	var blockMapType = BlockMapType{BlockMap: blockMap}
	var blocks []Block

	for _, tok := range proseDoc.Tokens() {
		// We want to stem and minify
		stem := nlp.Stemmer(tok.Text)
		miniStem := nlp.Minifier(stem)

		err = blockMapType.UpdateBlockMap(miniStem, tok)
	}

	// convert blockMap to array of blocks
	for _, block := range blockMap {
		blocks = append(blocks, block)
	}

	profile.Blocks = blocks
	profile.Name = rawDoc.Key
	profile.Tag = rawDoc.Tag

	return profile, nil
}
