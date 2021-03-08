package nlp

import (
	"testing"
)

func TestNormalizeText(t *testing.T) {

	var testStringsToNormalize = []struct {
		tname    string
		strTest  string
		expected string
	}{
		{
			"Test One for Ascii Fold and Remove Punctuation",
			"<b>@!Is history the historical déja vu *()hysterical?</b>\n",
			"Is history the historical deja vu hysterical",
		},
		{
			"Test elNino1 for AsciiFold",
			"ElNin\u0303o",
			"ElNino",
		},
	}
	for _, test := range testStringsToNormalize {
		if got, _ := NormalizeText(test.strTest); &got != nil && got != test.expected {
			t.Errorf("Failed! %s : \n%s \ndoes not match expected : \n%s\n", test.tname, test.strTest, test.expected)
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
			"LOAN OFFICERAccountant",
			"LOAN OFFICER Accountant",
		},
		{
			"Test Squooshed Around Comma for Splitting",
			"answer the phone,with",
			"answer the phone, with",
		},
	}
	for _, test := range testStringsToSplit {
		if got := CamelCaseSplitter(test.strTest); &got != nil && got != test.expected {
			t.Errorf("Failed! : \n%s \n%s does not match expected : \n%s\n", test.tname, got, test.expected)
		}
	}
}
