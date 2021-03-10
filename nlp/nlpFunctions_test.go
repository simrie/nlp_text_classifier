package nlp

import (
	"testing"

	"github.com/jdkato/prose/v3"
)

func TestNlpFunctions_MakeSegmenter(t *testing.T) {

	var testName string = "TestMakeSegmenter returns a *prose.Document"
	var text string = "Baby Chickens"

	//var token1 prose.Token = prose.Token{Tag: "NN", Text: "Baby"}
	//var token2 prose.Token = prose.Token{Tag: "NS", Text: "Chickens"}

	var got interface{}
	var err error

	got, err = MakeSegmenter(text)
	if err != nil {
		t.Errorf("Failed! %s : \nreturned an error %v\n", testName, err)
	}
	_, ok := got.(*prose.Document)
	if !ok {
		t.Errorf("Failed! %s : \nreturned not of type prose.Document %v\n", testName, got)
	}
}

func TestNlpFunctions_Stemmer(t *testing.T) {

	var testStrings = []struct {
		tname    string
		strTest  string
		expected string
	}{
		{
			"Test Stemmer 1",
			"Historical",
			"histor",
		},
		{
			"Test Stemmer 2",
			"Baby",
			"babi",
		},
		{
			"Test Stemmer 3",
			"Babies",
			"babi",
		},
	}
	for _, test := range testStrings {
		if got := Stemmer(test.strTest); &got != nil && got != test.expected {
			t.Errorf("Failed! %s : \n%s \ndoes not match expected : \n%s\n", test.tname, got, test.expected)
		}
	}
}

func TestNlpFunctions_PreserveMinPrefix(t *testing.T) {

	var testStrings = []struct {
		tname     string
		strTest   string
		expected1 string
		expected2 string
	}{
		{
			"Test PreserveMinPrefix 1",
			"Historical",
			"Hist",
			"orical",
		},
		{
			"Test PreserveMinPrefix 2",
			"aardvark",
			"aard",
			"vark",
		},
	}
	for _, test := range testStrings {
		if got1, got2 := PreserveMinPrefix(test.strTest, 4); got1 == "" || got1 != test.expected1 || got2 != test.expected2 {
			t.Errorf("Failed! %s : \n%s is blank or does not match expected %s, or \n%s does not match expected %s\n", test.tname, got1, test.expected1, got2, test.expected2)
		}
	}
}

func TestNlpFunctions_Minifier(t *testing.T) {

	stem1 := Stemmer("Historical")
	stem2 := Stemmer("Chickens")
	stem3 := Stemmer("supercalifragilisticexpialidocious")

	var testStrings = []struct {
		tname    string
		strTest  string
		expected string
	}{
		{
			"Test Minifier 1",
			stem1,
			"HISTR",
		},
		{
			"Test Minifier 2",
			stem2,
			"CHICKN",
		},
		{
			"Test Minifier 3",
			stem3,
			"SUPERCLFRGLSTCXPLDC",
		},
	}
	for _, test := range testStrings {
		if got := Minifier(test.strTest); &got != nil && got != test.expected {
			t.Errorf("Failed! %s : \n%s \ndoes not match expected : \n%s\n", test.tname, got, test.expected)
		}
	}
}
