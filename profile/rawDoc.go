package profile

import (
	"nlp_text_classifier/nlp"
	"nlp_text_classifier/utils"

	"github.com/jdkato/prose/v3"
)

/*
RawDoc holds incoming doc objects
*/
type RawDoc struct {
	Key  string `json:"key"`
	Text string `json:"text"`
	Tag  string `json:"tag"`
}

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

	acceptableTokenTypeMap := utils.MapStringSlice(nlp.ProseNounsVerbsAdjAdv)

	for _, tok := range proseDoc.Tokens() {
		// Do not process tokens if these are not of a type we want
		_, ok := acceptableTokenTypeMap[tok.Tag]
		if !ok {
			continue
		}
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
