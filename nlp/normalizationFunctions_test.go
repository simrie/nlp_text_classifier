package nlp

import (
	"testing"
)

func TestAsciiFold(t *testing.T) {

	var testStringsToNormalize = []struct {
		tname    string
		strTest  string
		expected string
	}{
		{
			"Test Ascii Fold and Remove Punctuation",
			"<b>@!This déja vu *()mythical?</b>\n",
			"This deja vu mythical",
		},
		{
			"Test elNino1 for AsciiFold",
			"ElNin\u0303o",
			"ElNino",
		},
		{
			"Test resume for AsciiFold",
			"résumé",
			"resume",
		},
	}
	for _, test := range testStringsToNormalize {
		if got, _ := NormalizeText(test.strTest); &got != nil && got != test.expected {
			t.Errorf("Failed! %s : \n%s \ndoes not match expected : \n%s\n", test.tname, got, test.expected)
		}
	}
}

func TestRemoveHTMLTags(t *testing.T) {

	var testStringsToNormalize = []struct {
		tname    string
		strTest  string
		expected string
	}{
		{
			"Test RemoveHTMLTags Lists mismatched",
			"<ul><li>item 1</li><li>item 2</li></ol>",
			"item 1 item 2",
		},
		{
			"Test RemoveHTMLTags Various",
			"<b>Hey! <br>New Line <p>Paragraph</b>",
			"Hey New Line Paragraph",
		},
	}
	for _, test := range testStringsToNormalize {
		if got, _ := NormalizeText(test.strTest); &got != nil && got != test.expected {
			t.Errorf("Failed! %s : \n%s \ndoes not match expected : \n%s\n", test.tname, got, test.expected)
		}
	}
}

func TestRemoveTabsAndLineFeeds(t *testing.T) {

	var testStringsToNormalize = []struct {
		tname    string
		strTest  string
		expected string
	}{
		{
			"Test RemoveTabsAndLineFeeds",
			"Hello \t there \nyou\nlook\nnice today",
			"Hello there you look nice today",
		},
		{
			"Test RemoveTabsAndLineFeeds Various",
			"\n\t1\t2\t3\t4\n\n",
			"1 2 3 4",
		},
	}
	for _, test := range testStringsToNormalize {
		if got, _ := NormalizeText(test.strTest); &got != nil && got != test.expected {
			t.Errorf("Failed! %s : \n%s \ndoes not match expected : \n%s\n", test.tname, got, test.expected)
		}
	}
}

func TestRemoveNonSpaceNonAlphanumeric(t *testing.T) {

	var testStringsToNormalize = []struct {
		tname    string
		strTest  string
		expected string
	}{
		{
			"Test RemoveNonSpaceNonAlphanumeric punctuation",
			"Hello@you.com, #whatever ^:-)",
			"Hello you com whatever",
		},
		{
			"Test RemoveNonSpaceNonAlphanumeric symbols",
			"The £ pound is not € a euro",
			"The pound is not a euro",
		},
	}
	for _, test := range testStringsToNormalize {
		if got, _ := NormalizeText(test.strTest); &got != nil && got != test.expected {
			t.Errorf("Failed! %s : \n%s \ndoes not match expected : \n%s\n", test.tname, got, test.expected)
		}
	}
}

func TestCamelCaseSplitting(t *testing.T) {

	var testStringsToSplit = []struct {
		tname    string
		strTest  string
		expected string
	}{
		{
			"Test Pushed Together Words for Splitting",
			"NextEra Energy",
			"Next Era Energy",
		},
		{
			"Test Comma-squooshed text for Splitting",
			"coffee,with cream",
			"coffee, with cream",
		},
	}
	for _, test := range testStringsToSplit {
		if got := CamelCaseSplitter(test.strTest); &got != nil && got != test.expected {
			t.Errorf("Failed! : \n%s \n%s does not match expected : \n%s\n", test.tname, got, test.expected)
		}
	}
}
