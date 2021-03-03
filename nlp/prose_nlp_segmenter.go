package nlp

import (
	"fmt"
	porterStemmer "github.com/agonopol/go-stem"
	"github.com/jdkato/prose/v3"
	"strings"
)

func MakeSegmenter(text string) (*prose.Document, error) {
	return prose.NewDocument(
		text,
		prose.WithTokenization(true),
		prose.WithTagging(false),
		prose.WithExtraction(true))
}

func Stemmer(text string) string {
	word := []byte(text)
	stem := porterStemmer.Stem(word)
	return string(stem[:])
}

func PreserveMinPrefix(text string, min_length int) (string, string) {
	var actual_length = len(text)
	var keep_prefix = ""
	var rest_of_word = ""
	if actual_length <= min_length {
		keep_prefix = text
		return keep_prefix, rest_of_word
	}
	keep_prefix = text[0:min_length]
	rest_of_word = text[min_length:actual_length]
	return keep_prefix, rest_of_word
}

func Minifier(text string) string {
	var vowels = [5]string{"a", "e", "i", "o", "u"}
	var word string = text
	word = strings.ToLower(word)
	word = strings.TrimSpace(word)

	// keep all if length less than
	var minPrefix = 4
	var keepPrefix, restOfWord = PreserveMinPrefix(word, minPrefix)

	// preserve first letter if vowel
	//var keepVowel = PreserveInitialVowel(word)

	// remove vowels
	for _, vowel := range vowels {
		restOfWord = strings.ReplaceAll(restOfWord, vowel, "")
	}

	word = fmt.Sprintf("%s%s", keepPrefix, restOfWord)
	//word = fmt.Sprintf("%s%s", keepVowel, word)
	// convert to caps
	word = strings.ToUpper(word)

	// remove double-consonants

	return word
}
