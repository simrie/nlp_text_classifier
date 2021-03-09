package nlp

import (
	"regexp"
	"sort"
	"strings"
	"unicode"

	camelCaseSplitter "github.com/fatih/camelcase"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

/*
ASCIIFold removes accents from strings
//https://stackoverflow.com/questions/24588295/go-removing-accents-from-strings
*/
func ASCIIFold(in string) string {
	if in == "" {
		return in
	}
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	normStr1, _, _ := transform.String(t, in)
	return normStr1
}

/*
RemoveHTMLTags removes tags leaving content
*/
func RemoveHTMLTags(in string) string {
	if in == "" {
		return in
	}
	// match html tag and replace it with ""
	// regex to match html tag
	const pattern = `(<\/?[a-zA-Z]+?[^>]*\/?>)*`
	r := regexp.MustCompile(pattern)
	groups := r.FindAllString(in, -1)

	// should replace long string first
	sort.Slice(groups, func(i, j int) bool {
		return len(groups[i]) > len(groups[j])
	})

	for _, group := range groups {
		if group == "" {
			continue
		}
		lcGroup := strings.ToLower(strings.TrimSpace(group))
		if lcGroup != "" {
			in = strings.ReplaceAll(in, group, " ")
			continue
		}
	}
	return in
}

/*
RemoveTabsAndLineFeeds removes tabs and line feeds
*/
func RemoveTabsAndLineFeeds(in string) string {
	if in == "" {
		return in
	}
	in = strings.ReplaceAll(in, "\n", " ")
	in = strings.ReplaceAll(in, "\t", " ")
	in = strings.TrimSpace(in)
	return in
}

/*
RemoveNonSpaceNonAlphanumeric removes all but letters and numbers
*/
func RemoveNonSpaceNonAlphanumeric(in string) (string, error) {
	if in == "" {
		return in, nil
	}
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		return in, err
	}
	processedString := reg.ReplaceAllString(in, " ")
	return processedString, nil
}

/*
CompressSpaces removes excessive spaces
*/
func CompressSpaces(in string) string {
	str := in
	str = strings.ReplaceAll(str, "   ", " ")
	str = strings.ReplaceAll(str, "  ", " ")
	return str
}

/*
CamelCaseSplitter splits camel case words
*/
func CamelCaseSplitter(text string) string {
	splitted := camelCaseSplitter.Split(text)
	if splitted[0] == text {
		return text
	}
	// Need special consideration for splitted punctuation
	// If apostrophe before or after word, do not join with space
	testJoin := strings.Join(splitted, "")
	words := ""
	if strings.Index(testJoin, "'") >= 0 {
		words = testJoin
	} else {
		// join back with spaces between
		words = strings.Join(splitted, " ")
	}
	// compensate for when splits were done on punctuation
	words = strings.ReplaceAll(words, " . ", ". ")
	words = strings.ReplaceAll(words, " , ", ", ")
	words = strings.ReplaceAll(words, " ; ", "; ")
	words = strings.ReplaceAll(words, " ! ", "! ")
	words = strings.ReplaceAll(words, " - ", "-")
	words = strings.ReplaceAll(words, "   ", " ")
	words = strings.ReplaceAll(words, "  ", " ")
	return words
}

/*
CamelCaseSplitAndRejoin returns split words joined by a string
*/
func CamelCaseSplitAndRejoin(text string) string {
	// TODO: call this as a Goroutine
	//       or not as part of text normalization
	//       because Prose Segmenter is a bottleneck in processing
	doc, _ := MakeSegmenter(text)
	var replacementList = make(map[string]string)

	for _, tok := range doc.Tokens() {
		splitted := CamelCaseSplitter(tok.Text)
		if splitted != tok.Text {
			replacementList[tok.Text] = splitted
		}
	}
	// apply replacements
	for key, replacement := range replacementList {
		text = strings.Replace(text, key, replacement, -1)
	}
	return text
}

/*
NormalizeText returns plain(er) text that nlp library can more easily tokenize
*/
func NormalizeText(str string) (string, error) {
	// TODO: Define flags for which normalization functions to apply, i.e.
	// "plain text" would use the functions shown here but not remove HTML tags
	// "html text" would strip the HTML tags and perhaps replace some entities with words
	// "camelCase" would first split any camelCase words so each word could be stemmed

	if str == "" {
		return str, nil
	}
	var err error
	str = ASCIIFold(str)
	str = RemoveHTMLTags(str)
	str = RemoveTabsAndLineFeeds(str)
	str, err = RemoveNonSpaceNonAlphanumeric(str)
	str = CompressSpaces(str)
	if err != nil {
		return "", err
	}
	//str = CamelCaseSplitAndRejoin(str)
	str = strings.TrimSpace(str)
	return str, nil
}
