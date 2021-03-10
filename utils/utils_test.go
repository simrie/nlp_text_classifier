package utils

import (
	"reflect"
	"testing"
)

func TestUtils_MapStringSlice(t *testing.T) {

	// converts a slice of words to map[string]{}

	var wordSlice1 []string
	wordSlice1 = []string{"cookies", "cream"}
	wordMap1 := make(map[string]struct{})
	wordMap1["cookies"] = struct{}{}
	wordMap1["cream"] = struct{}{}

	var wordSlice2 []string
	wordSlice2 = []string{"cookies", "cream", "chocolate", "donuts"}
	wordMap2 := make(map[string]struct{})
	wordMap2["cookies"] = struct{}{}
	wordMap2["cream"] = struct{}{}
	wordMap2["chocolate"] = struct{}{}
	wordMap2["donuts"] = struct{}{}

	var testUtils = []struct {
		tname    string
		sliceArg []string
		expected map[string]struct{}
	}{
		{
			"Test Utils MapStringSlice 1",
			wordSlice1,
			wordMap1,
		},
		{
			"Test Utils MapStringSlice 2",
			wordSlice2,
			wordMap2,
		},
	}
	for _, test := range testUtils {
		if got := MapStringSlice(test.sliceArg); &got != nil && !reflect.DeepEqual(got, test.expected) {
			t.Errorf("Failed! %s : \n%v \ndoes not match expected : \n%v\n", test.tname, got, test.expected)
		}
	}
}

func TestUtils_ContainsString(t *testing.T) {

	// true if string slice contains a string, else false

	var wordSlice []string
	wordSlice = []string{"cookies", "cream", "chocolate", "donuts"}

	wordTest1 := "cookies"
	bContains1 := true

	wordTest2 := "vanilla"
	bContains2 := false

	var testUtils = []struct {
		tname    string
		sliceArg []string
		itemArg  string
		expected bool
	}{
		{
			"Test Utils ContainsString 1",
			wordSlice,
			wordTest1,
			bContains1,
		},
		{
			"Test Utils ContainsString 2",
			wordSlice,
			wordTest2,
			bContains2,
		},
	}
	for _, test := range testUtils {
		if got := ContainsString(test.sliceArg, test.itemArg); &got == nil || got != test.expected {
			t.Errorf("Failed! %s : \n%v \nis nil or != expected : \n%v\n", test.tname, got, test.expected)
		}
	}
}

func TestUtils_KeysInCommon(t *testing.T) {

	// returns an array of strings common to two string arrays

	var wordSlice1 []string = []string{"cookies", "cream", "chocolate", "donuts"}
	var wordSlice2 []string = []string{"cookies", "cream", "vanilla", "cupcakes"}
	var wordSlice3 []string = []string{"cookies", "cream"}

	var testUtils = []struct {
		tname    string
		argA     []string
		argB     []string
		expected []string
	}{
		{
			"Test Utils KeysInCommon",
			wordSlice1,
			wordSlice2,
			wordSlice3,
		},
	}
	for _, test := range testUtils {
		if got := KeysInCommon(test.argA, test.argB); &got == nil || !reflect.DeepEqual(got, test.expected) {
			t.Errorf("Failed! %s : \n%v \nis nil or != expected : \n%v\n", test.tname, got, test.expected)
		}
	}
}

func TestUtils_KeysDiff(t *testing.T) {

	// pass two arrays of strings
	// first result has items only in first array
	// second result has items only in second array

	var wordSlice1 []string = []string{"cookies", "cream", "chocolate", "donuts"}
	var wordSlice2 []string = []string{"cookies", "cream", "vanilla", "cupcakes"}
	var wordSlice3 []string = []string{"chocolate", "donuts"}
	var wordSlice4 []string = []string{"vanilla", "cupcakes"}

	var testUtils = []struct {
		tname     string
		argA      []string
		argB      []string
		expected1 []string
		expected2 []string
	}{
		{
			"Test Utils KeysDiff",
			wordSlice1,
			wordSlice2,
			wordSlice3,
			wordSlice4,
		},
	}
	for _, test := range testUtils {
		if got1, got2 := KeysDiff(test.argA, test.argB); got1 == nil || got2 == nil || !reflect.DeepEqual(got1, test.expected1) || !reflect.DeepEqual(got2, test.expected2) {
			t.Errorf("Failed! %s : \n%v is nil or != expected1 :%v\nor %v is nil or != expected2 %v ", test.tname, got1, test.expected1, got2, test.expected2)
		}
	}
}
