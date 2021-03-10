package profile

import (
	"nlp_text_classifier/utils"
)

/*
WordSeen shows original words that were stemmed and their counts
*/
type WordSeen struct {
	Word string `json:"word"`
	Seen int    `json:"seen"`
}

/*
WordsSeenType is WordsSeen slice can be used as a method receiver
*/
type WordsSeenType struct {
	WordsSeen []WordSeen `json:"words_seen,omitempty"`
}

/*
WordsSeenMapType lets us easily look up words in WordsSeen
*/
type WordsSeenMapType struct {
	WordsSeenMap map[string]WordSeen `json:"words_seen_map,omitempty"`
}

/*
MapWordsSeen returns WordsSeenMapType containing WordsSeenMap
*/
func (wordsSeen *WordsSeenType) MapWordsSeen() WordsSeenMapType {
	var wordsSeenMapType WordsSeenMapType
	wordsSeenMap := make(map[string]WordSeen)
	for _, v := range wordsSeen.WordsSeen {
		wordsSeenMap[v.Word] = v
	}
	wordsSeenMapType.WordsSeenMap = wordsSeenMap
	return wordsSeenMapType
}

/*
MapKeysAsStrings returns slice of string keys from WordsSeenMapType.WordsSeenMap
*/
func (wordsSeenMapType *WordsSeenMapType) MapKeysAsStrings() []string {
	wordsSeenMap := wordsSeenMapType.WordsSeenMap
	keys := make([]string, len(wordsSeenMap))
	i := 0
	for k := range wordsSeenMap {
		keys[i] = k
		i++
	}
	return keys
}

/*
Combine []WordSeen array with info from incoming array
*/
func (wordSeen *WordSeen) Combine(inWordSeen WordSeen) error {
	if wordSeen.Word == inWordSeen.Word {
		wordSeen.Seen = wordSeen.Seen + inWordSeen.Seen
	}
	return nil
}

/*
Combine WordsSeenType.WordsSeen with new WordsSeen
*/
func (wordsSeen *WordsSeenType) Combine(inWordsSeen WordsSeenType) error {
	wordsSeenMapType := wordsSeen.MapWordsSeen()
	wordsSeenMap := wordsSeenMapType.WordsSeenMap
	wordsSeenKeys := wordsSeenMapType.MapKeysAsStrings()
	inWordsSeenMapType := inWordsSeen.MapWordsSeen()
	inWordsSeenMap := inWordsSeenMapType.WordsSeenMap
	inWordsSeenKeys := inWordsSeenMapType.MapKeysAsStrings()

	// compare Key slices

	//  find keys in common
	//  for keys in common between wordsSeenKeys and inWordsSeenKeys
	//  add the "Seen" count to wordsSeen
	var keysInCommon []string = utils.KeysInCommon(wordsSeenKeys, inWordsSeenKeys)
	var newWordsSeen []WordSeen = []WordSeen{}

	for _, key := range keysInCommon {
		wordSeen, _ := wordsSeenMap[key]
		inWordSeen, ok := inWordsSeenMap[key]
		if ok {
			wordSeen.Combine(inWordSeen)
		}
		newWordsSeen = append(newWordsSeen, wordSeen)
	}

	// Find keys to add from inWordsSeenKeys
	// if inWordsSeen has additional WordsSeen, append those to wordsSeen
	keysToAdd, keysToAdd2 := utils.KeysDiff(wordsSeenKeys, inWordsSeenKeys)
	keysToAdd = append(keysToAdd, keysToAdd2...)

	for _, key := range keysToAdd {
		wordSeen, ok := wordsSeenMap[key]
		if ok {
			newWordsSeen = append(newWordsSeen, wordSeen)
			continue
		}
		inWordSeen, ok := inWordsSeenMap[key]
		if ok {
			newWordsSeen = append(newWordsSeen, inWordSeen)
		}
	}

	wordsSeen.WordsSeen = newWordsSeen
	return nil
}
