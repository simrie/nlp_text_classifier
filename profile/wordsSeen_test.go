package profile

import (
	"reflect"
	"testing"
)

func TestWordsSeen_MapWordsSeen(t *testing.T) {

	// converts WordsSeenType to WordsSeenMapType

	wordSeen1 := WordSeen{Word: "HELLO", Seen: 1}
	wordSeen2 := WordSeen{Word: "GOODBYE", Seen: 1}
	wordsSeen := []WordSeen{wordSeen1, wordSeen2}

	var wordsSeenType WordsSeenType
	wordsSeenType.WordsSeen = wordsSeen

	var wordsSeenMapType WordsSeenMapType
	wordsSeenMap := make(map[string]WordSeen)
	wordsSeenMap["HELLO"] = wordSeen1
	wordsSeenMap["GOODBYE"] = wordSeen2

	wordsSeenMapType.WordsSeenMap = wordsSeenMap

	var testWordsSeen = []struct {
		tname    string
		expected WordsSeenMapType
	}{
		{
			"Test WordsSeen MapWordsSeen",
			wordsSeenMapType,
		},
	}
	for _, test := range testWordsSeen {
		if got := wordsSeenType.MapWordsSeen(); &got != nil && !reflect.DeepEqual(got, test.expected) {
			t.Errorf("Failed! %s : \n%v \ndoes not match expected : \n%v\n", test.tname, got, test.expected)
		}
	}
}

func TestWordsSeen_MapKeysAsStrings(t *testing.T) {

	// converts WordsSeenMapType to a slice of Strings

	var expectedSlice []string = []string{"HELLO", "GOODBYE"}

	var wordsSeenMapType WordsSeenMapType
	wordsSeenMap := make(map[string]WordSeen)
	wordsSeenMap["HELLO"] = WordSeen{Word: "HELLO", Seen: 1}
	wordsSeenMap["GOODBYE"] = WordSeen{Word: "GOODBYE", Seen: 1}

	wordsSeenMapType.WordsSeenMap = wordsSeenMap

	var testWordsSeen = []struct {
		tname    string
		expected []string
	}{
		{
			"Test WordsSeen MapKeysAsStrings",
			expectedSlice,
		},
	}
	for _, test := range testWordsSeen {
		if got := wordsSeenMapType.MapKeysAsStrings(); &got != nil && !reflect.DeepEqual(got, test.expected) {
			t.Errorf("Failed! %s : \n%v \ndoes not match expected : \n%v\n", test.tname, got, test.expected)
		}
	}
}

func TestWordsSeen_CombineWordSeen(t *testing.T) {

	// combines Seen count of two WordsSeen with same Word
	wordSeen1 := WordSeen{Word: "HELLO", Seen: 2}
	wordSeen2 := WordSeen{Word: "HELLO", Seen: 3}
	wordSeen3 := WordSeen{Word: "HELLO", Seen: 5}

	var pWordSeen1 *WordSeen = &wordSeen1
	var pWordSeen2 *WordSeen = &wordSeen2

	var testWordsSeen = []struct {
		tname    string
		receiver *WordSeen
		argument WordSeen
		expected WordSeen
	}{
		{
			"Test WordSeen Combine 1",
			pWordSeen1,
			wordSeen2,
			wordSeen3,
		},
		{
			"Test WordSeen Combine 2",
			pWordSeen2,
			wordSeen1,
			wordSeen3,
		},
	}
	for _, test := range testWordsSeen {
		if got := test.receiver.Combine(test.argument); got != nil || !reflect.DeepEqual(*test.receiver, test.expected) {
			t.Errorf("Failed! %s : \nerr: %v \nOR %v does not match expected : \n%v\n", test.tname, got, test.receiver, test.expected)
		}
	}
}

func TestWordsSeen_CombineWordSeenType(t *testing.T) {

	// combines wordsSeen from another WordsSeenType
	wordSeen1 := WordSeen{Word: "HAM", Seen: 2}
	wordSeen2 := WordSeen{Word: "EGG", Seen: 2}
	wordSeen3 := WordSeen{Word: "CHEESE", Seen: 2}
	wordSeen4 := WordSeen{Word: "BACON", Seen: 3}
	wordSeen5 := WordSeen{Word: "EGG", Seen: 3}
	wordSeen6 := WordSeen{Word: "CHEESE", Seen: 3}
	wordSeen7 := WordSeen{Word: "EGG", Seen: 5}
	wordSeen8 := WordSeen{Word: "CHEESE", Seen: 5}

	var wordsSeenType1 WordsSeenType
	wordsSeenType1.WordsSeen = []WordSeen{wordSeen1, wordSeen2, wordSeen3}

	var wordsSeenType2 WordsSeenType
	wordsSeenType2.WordsSeen = []WordSeen{wordSeen4, wordSeen5, wordSeen6}

	var wordsSeenType3 WordsSeenType
	// words in common + words from base + words from added
	// egg, cheese, ham, bacon
	wordsSeenType3.WordsSeen = []WordSeen{wordSeen7, wordSeen8, wordSeen1, wordSeen4}

	var testWordsSeen = []struct {
		tname    string
		receiver WordsSeenType
		argument WordsSeenType
		expected WordsSeenType
	}{
		{
			"Test WordsSeenType Combine",
			wordsSeenType1,
			wordsSeenType2,
			wordsSeenType3,
		},
	}
	for _, test := range testWordsSeen {
		if got := test.receiver.Combine(test.argument); got != nil || !reflect.DeepEqual(test.receiver, test.expected) {
			t.Errorf("Failed! %s : \nerr: %v \nOR %v does not match expected : \n%v\n", test.tname, got, test.receiver, test.expected)
		}
	}
}
