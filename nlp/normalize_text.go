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

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func AsciiFold(in string) string {
	if in == "" {
		return in
	}
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	normStr1, _, _ := transform.String(t, in)
	return normStr1
}

func FixBrokenHtmlEntities(in string) string {
	//&apos;
	apos := " & apos; s"
	in = strings.ReplaceAll(in, apos, "'s")
	apos = " & apos;s"
	in = strings.ReplaceAll(in, apos, "'s")
	apos = "&apos;s"
	in = strings.ReplaceAll(in, apos, "'s")
	apos = "& apos;s"
	in = strings.ReplaceAll(in, apos, "'s")
	apos = "& apos;"
	in = strings.ReplaceAll(in, apos, "'")
	apos = "&apos;"
	in = strings.ReplaceAll(in, apos, "'")
	apos = " &apos;"
	in = strings.ReplaceAll(in, apos, "'")
	//&amp;
	amp := "& amp;"
	in = strings.ReplaceAll(in, amp, "&")
	amp = "&amp;"
	in = strings.ReplaceAll(in, amp, "&")
	//&nbsp;
	nbsp := "& nbsp;"
	in = strings.ReplaceAll(in, nbsp, " ")
	nbsp = "&nbsp;"
	in = strings.ReplaceAll(in, nbsp, " ")
	//&lt;
	lt := "& lt;"
	in = strings.ReplaceAll(in, lt, "<")
	lt = "&lt;"
	in = strings.ReplaceAll(in, lt, "<")
	//&gt;
	gt := "& gt;"
	in = strings.ReplaceAll(in, gt, ">")
	gt = "&gt;"
	in = strings.ReplaceAll(in, gt, ">")
	//&quot;
	quot := "& quot;"
	in = strings.ReplaceAll(in, quot, "'")
	quot = "&quot;"
	in = strings.ReplaceAll(in, quot, "'")
	//&reg; [TM], &circledR;
	reg := "& reg;"
	in = strings.ReplaceAll(in, reg, "[TM]")
	reg = "&reg;"
	in = strings.ReplaceAll(in, reg, "[TM]")
	reg = "& circledR;"
	in = strings.ReplaceAll(in, reg, "[TM]")
	reg = "&circledR;"
	in = strings.ReplaceAll(in, reg, "[TM]")
	//&copy; [copyright]
	copy := "& copy;"
	in = strings.ReplaceAll(in, copy, "[copyright]")
	copy = "&copy;"
	in = strings.ReplaceAll(in, copy, "[copyright]")
	//&cent; [cents]
	cent := "& cent;"
	in = strings.ReplaceAll(in, cent, "[cents]")
	cent = "&cent;"
	in = strings.ReplaceAll(in, cent, "[cents]")
	//&pound; [pounds]
	pound := "& pound;"
	in = strings.ReplaceAll(in, pound, "[pounds]")
	pound = "&pound;"
	in = strings.ReplaceAll(in, pound, "[pounds]")
	//&yen; [yen]
	yen := "& yen;"
	in = strings.ReplaceAll(in, yen, "[yen]")
	yen = "&yen;"
	in = strings.ReplaceAll(in, yen, "[yen]")
	//&euro; [euros]
	euro := "& euro;"
	in = strings.ReplaceAll(in, euro, "[euros]")
	euro = "&euro;"
	in = strings.ReplaceAll(in, euro, "[euros]")
	//&commat;
	commat := "& commat;"
	in = strings.ReplaceAll(in, commat, "@")
	commat = "&commat;"
	in = strings.ReplaceAll(in, commat, "@")
	//&bull;
	bull := "& bull;"
	in = strings.ReplaceAll(in, bull, " *")
	bull = "&bull;"
	in = strings.ReplaceAll(in, bull, " *")
	return in
}

func ReplaceHtmlListItemsWithColon(in string) string {
	if in == "" {
		return in
	}
	// match html tag and replace it with ""
	// regex to match html tag
	const pattern = `(<\/?[a-zA-Z]+?[^>]*\/?>)*`
	r := regexp.MustCompile(pattern)
	groups := r.FindAllString(in, -1)

	// sort shorter items to replace list items with colon
	sort.Slice(groups, func(i, j int) bool {
		return len(groups[i]) < len(groups[j])
	})
	for _, group := range groups {
		if group == "" {
			continue
		}
		lc_group := strings.ToLower(strings.TrimSpace(group))
		if lc_group == "<br>" {
			in = strings.ReplaceAll(in, group, " ")
			continue
		}
		if lc_group == "<ul><li>" {
			in = strings.ReplaceAll(in, group, " ")
			continue
		}
		if lc_group == "</li><li>" {
			in = strings.ReplaceAll(in, group, "; ")
			continue
		}
		if lc_group == "</li></ul>" {
			in = strings.ReplaceAll(in, group, ". ")
			continue
		}
		if lc_group == "<ul><li>" {
			in = strings.ReplaceAll(in, group, " ")
			continue
		}
	}
	// fix ".;"
	in = strings.ReplaceAll(in, ".;", ". ")
	return in
}

func RemoveHtmlTag(in string) string {
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
		lc_group := strings.ToLower(strings.TrimSpace(group))
		if lc_group != "" {
			in = strings.ReplaceAll(in, group, "")
			continue
		}
	}
	return in
}

func RemoveTabsAndLineFeeds(in string) string {
	if in == "" {
		return in
	}
	in = strings.ReplaceAll(in, "\n", " ")
	in = strings.ReplaceAll(in, "\t", " ")
	in = strings.TrimSpace(in)
	return in
}

func RemoveNonSpaceNonAlphanumeric(in string) (string, error) {
	if in == "" {
		return in, nil
	}
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		return in, err
	}
	processedString := reg.ReplaceAllString(in, "")
	return processedString, nil
}

func CompressSpaces(in string) string {
	str := in
	str = strings.ReplaceAll(str, "   ", " ")
	str = strings.ReplaceAll(str, "  ", " ")
	return str
}

func CamelCaseSplitter(text string) string {
	splitted := camelCaseSplitter.Split(text)
	if splitted[0] == text {
		return text
	}
	// Need special consideration for splitted punctuation
	// If apostrophe before or after word, do not join with space
	test_join := strings.Join(splitted, "")
	words := ""
	if strings.Index(test_join, "'") >= 0 {
		words = test_join
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

func CamelCaseSplitAndRejoin(text string) string {
	doc, _ := MakeSegmenter(text)
	var replacementList = make(map[string]string)

	var valid = make(map[string]bool)
	for _, v := range ProseTagList_WordsOnly {
		valid[v] = true
	}

	for _, tok := range doc.Tokens() {

		//if StringInArrayUpper(tok.Tag, ProseTagList_WordsOnly) {
		if valid[tok.Tag] {
			splitted := CamelCaseSplitter(tok.Text)
			if splitted != tok.Text {
				var tokens []string
				tokens = append(tokens, splitted)
				split_joined := strings.Join(tokens, " ")
				replacementList[tok.Text] = split_joined
			}

		}
	}
	// apply replacements
	for key, replacement := range replacementList {
		text = strings.Replace(text, key, replacement, -1)
	}
	return text
}

func NormalizeText(str string) (string, error) {
	// this should already have been cleaned of anything that goes into an addback list
	if str == "" {
		return str, nil
	}
	var err error
	str = AsciiFold(str)
	str = RemoveHtmlTag(str)
	str = RemoveTabsAndLineFeeds(str)
	str, err = RemoveNonSpaceNonAlphanumeric(str)
	if err != nil {
		return "", err
	}
	//str = CamelCaseSplitAndRejoin(str)
	str = strings.TrimSpace(str)
	return str, nil
}
