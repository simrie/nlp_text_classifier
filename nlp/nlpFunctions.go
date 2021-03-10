package nlp

import (
	"fmt"
	"strings"

	porterStemmer "github.com/agonopol/go-stem"
	"github.com/jdkato/prose/v3"
)

/*
ProseTagListWordsOnly is an array of the Prose token types
for most nouns, pronouns, verbs, adjectives, adverbs, and the determiner "not"

These correspond to the Prose Document.Token part-of-speech tags.
see https://github.com/jdkato/prose#segmenting
*/

/*
IgnoreMinifiedList is an array of words that can be ignored
*/
var IgnoreMinifiedList = []string{"BE", "IS", "ISNT", "AR", "ARENT", "AINT", "HAVE", "AN", "A", "IF", "WILL"}

/*
ProseNouns is an array of Prose tokens that represent nouns
*/
var ProseNouns = []string{"NN", "NS", "NNP", "NNPS", "NNS"}

/*
ProseVerbs is an array of Prose tokens that represent verbs
*/
var ProseVerbs = []string{"VB", "VBD", "VBG", "VBG", "VBN", "VBP", "VBZ"}

/*
ProseAdjectives is an array of Prose tokens that represent Adjectives
*/
var ProseAdjectives = []string{"JJ", "JJR", "JJS"}

/*
ProseAdverbs is an array of Prose tokens that represent Adverbs
*/
var ProseAdverbs = []string{"RB", "RBR", "RBS", "RP"}

/*
ProseDeterminers is an array of Prose tokens that represent Determiners
*/
var ProseDeterminers = []string{"DT"}

/*
ProseNounsVerbs is an array of Prose tokens for Nounse and Verbs
*/
var ProseNounsVerbs = append(ProseNouns, ProseVerbs...)

/*
ProseNounsVerbsAdverbs is an array of Prose tokens for nouns, verbs and adverbs
*/
var ProseNounsVerbsAdverbs = append(ProseNounsVerbs, ProseAdverbs...)

/*
ProseAdjectivesAdverbs is an array of Prose tokens for adjectives and adverbs
*/
var ProseAdjectivesAdverbs = append(ProseAdjectives, ProseAdverbs...)

/*
ProseNounsVerbsAdjAdv is an array of Prose tokens for nouns, verbs, adjectives and adverbs
*/
var ProseNounsVerbsAdjAdv = append(ProseNounsVerbs, ProseAdjectivesAdverbs...)

/*
ProseAdjAdvWithDeterminers is an array of Prose tokens for nouns, verbs, adjectives, adverbs and determiners
*/
var ProseAdjAdvWithDeterminers = append(ProseAdjectivesAdverbs, ProseDeterminers...)

/*
ProseNounsVerbsAdjAdvWithDeterminers is an array of Prose tokens for nouns, verbs, adjectives, adverbs and determiners
*/
var ProseNounsVerbsAdjAdvWithDeterminers = append(ProseNounsVerbsAdjAdv, ProseDeterminers...)

/*
MakeSegmenter converts text to a prose.Document consisting of Prose.tokens
*/
func MakeSegmenter(text string) (*prose.Document, error) {
	return prose.NewDocument(
		text,
		prose.WithTokenization(true),
		prose.WithTagging(false),
		prose.WithExtraction(true))
}

/*
Stemmer returns the stem returned by the PorterStemmer
*/
func Stemmer(text string) string {
	word := []byte(text)
	stem := porterStemmer.Stem(word)
	return string(stem[:])
}

/*
PreserveMinPrefix returns the first minLength letters of a stem and the remainder
*/
func PreserveMinPrefix(text string, minLength int) (string, string) {
	var actualLength = len(text)
	var keep = ""
	var remainder = ""
	if actualLength <= minLength {
		keep = text
		return keep, remainder
	}
	keep = text[0:minLength]
	remainder = text[minLength:actualLength]
	return keep, remainder
}

/*
Minifier returns the stem with the first X letters preserved but for readability but trimmed of extra vowels at the end
*/
func Minifier(text string) string {
	var vowels = [5]string{"a", "e", "i", "o", "u"}
	var word string = text
	word = strings.ToLower(word)
	word = strings.TrimSpace(word)

	// keep all if length less than minPrefix
	var minPrefix = 4
	var keepPrefix, restOfWord = PreserveMinPrefix(word, minPrefix)

	// remove vowels
	for _, vowel := range vowels {
		restOfWord = strings.ReplaceAll(restOfWord, vowel, "")
	}

	word = fmt.Sprintf("%s%s", keepPrefix, restOfWord)
	word = strings.ToUpper(word)

	// TODO: consider singular consonent for double-consonents
	return word
}
