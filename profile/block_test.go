package profile

import (
	"reflect"
	"testing"

	"github.com/jdkato/prose/v3"
)

func TestBlock_MapKeysAsStrings(t *testing.T) {

	// converts BlockMapType to a slice of Strings

	var expectedSlice []string = []string{"RESPONSIB", "HISTOR"}

	wordSeen1 := WordSeen{Word: "responsible", Seen: 1}
	wordSeen2 := WordSeen{Word: "responsibility", Seen: 1}
	wordsSeenA := []WordSeen{wordSeen1, wordSeen2}

	wordSeen3 := WordSeen{Word: "history", Seen: 1}
	wordSeen4 := WordSeen{Word: "historical", Seen: 1}
	wordsSeenB := []WordSeen{wordSeen3, wordSeen4}

	var block1 Block
	block1.Sources = wordsSeenA
	block1.Weight = 1
	block1.Count = 2
	block1.MiniStem = "RESPNSIB"

	var block2 Block
	block2.Sources = wordsSeenB
	block2.Weight = 1
	block2.Count = 2
	block2.MiniStem = "HISTOR"

	blockMap := make(map[string]Block)
	blockMap["RESPONSIB"] = block1
	blockMap["HISTOR"] = block2
	var blockMapType BlockMapType
	blockMapType.BlockMap = blockMap

	var testBlockMapType = []struct {
		tname    string
		expected []string
	}{
		{
			"Test Block MapKeysAsStrings",
			expectedSlice,
		},
	}
	for _, test := range testBlockMapType {
		if got := blockMapType.MapKeysAsStrings(); &got != nil && !reflect.DeepEqual(got, test.expected) {
			t.Errorf("Failed! %s : \n%v \ndoes not match expected : \n%v\n", test.tname, got, test.expected)
		}
	}
}

func TestBlock_CombineBlockMapType(t *testing.T) {

	// combines blockMap entries from arrqy of blockMapType

	wordSeen1 := WordSeen{Word: "gardening", Seen: 1}
	wordSeen2 := WordSeen{Word: "garden", Seen: 1}
	wordSeen3 := WordSeen{Word: "gardener", Seen: 1}

	wordSeen4 := WordSeen{Word: "landscaping", Seen: 1}
	wordSeen5 := WordSeen{Word: "landscaper", Seen: 1}

	wordSeen6 := WordSeen{Word: "manager", Seen: 1}
	wordSeen7 := WordSeen{Word: "manage", Seen: 1}
	wordSeen8 := WordSeen{Word: "management", Seen: 1}

	wordSeen9 := WordSeen{Word: "excavator", Seen: 1}
	wordSeen10 := WordSeen{Word: "excavating", Seen: 1}

	var block1 Block
	block1.Count = 3
	block1.Weight = 1
	block1.MiniStem = "GARDN"
	block1.Sources = []WordSeen{wordSeen1, wordSeen2, wordSeen3}

	var blockMapType1 BlockMapType
	blockMap1 := make(map[string]Block)
	blockMap1["GARDN"] = block1
	blockMapType1.BlockMap = blockMap1

	var block2 Block
	block2.Count = 2
	block2.Weight = 1
	block2.MiniStem = "LANDSCP"
	block2.Sources = []WordSeen{wordSeen4, wordSeen5}

	var blockMapType2 BlockMapType
	blockMap2 := make(map[string]Block)
	blockMap2["LANDSCP"] = block2
	blockMapType2.BlockMap = blockMap2

	var block3 Block
	block3.Count = 3
	block3.Weight = 1
	block3.MiniStem = "MANAG"
	block3.Sources = []WordSeen{wordSeen6, wordSeen7, wordSeen8}

	var block4 Block
	block4.Count = 2
	block4.Weight = 1
	block4.MiniStem = "EXCAVT"
	block4.Sources = []WordSeen{wordSeen9, wordSeen10}

	var blockMapType3 BlockMapType
	blockMap3 := make(map[string]Block)
	blockMap3["MANAG"] = block3
	blockMap3["EXCAVT"] = block4
	blockMapType3.BlockMap = blockMap3

	blockMapTypes := []BlockMapType{blockMapType2, blockMapType3}

	var blockMapType4 BlockMapType
	blockMap4 := make(map[string]Block)
	blockMap4["GARDN"] = block1
	blockMap4["LANDSCP"] = block2
	blockMap4["MANAG"] = block3
	blockMap4["EXCAVT"] = block4
	blockMapType4.BlockMap = blockMap4

	var testBlockMapType = []struct {
		tname    string
		receiver *BlockMapType
		argument []BlockMapType
		expected *BlockMapType
	}{
		{
			"Test BlockMapType Combine",
			&blockMapType1,
			blockMapTypes,
			&blockMapType4,
		},
	}
	for _, test := range testBlockMapType {
		if got := test.receiver.Combine(test.argument); got != nil || !reflect.DeepEqual(test.receiver, test.expected) {
			t.Errorf("Failed! %s : \nerr: %v \nOR %v does not match expected : \n%v\n", test.tname, got, test.receiver, test.expected)
		}
	}
}

func TestBlock_UpdateBlockMapType(t *testing.T) {

	miniStem := "TEST"

	wordSeen1 := WordSeen{Word: "test", Seen: 1}
	wordSeen2 := WordSeen{Word: "testing", Seen: 1}
	// Expect this one to be added to the block map from the token
	wordSeen3 := WordSeen{Word: "tested", Seen: 1}

	var block1 Block
	block1.Count = 2
	block1.Weight = 1
	block1.MiniStem = miniStem
	block1.Sources = []WordSeen{wordSeen1, wordSeen2}

	blockMap1 := make(map[string]Block)
	blockMap1[miniStem] = block1

	var blockMapType1 BlockMapType
	blockMapType1.BlockMap = blockMap1

	var block2 Block
	block2.Count = 3
	block2.Weight = 1
	block2.MiniStem = miniStem
	block2.Sources = []WordSeen{wordSeen1, wordSeen2, wordSeen3}

	blockMap2 := make(map[string]Block)
	blockMap2[miniStem] = block2

	var blockMapType2 BlockMapType
	blockMapType2.BlockMap = blockMap2

	var token prose.Token = prose.Token{Tag: "ADJ", Text: "tested"}

	var testBlockMapType = []struct {
		tname    string
		receiver *BlockMapType
		miniStem string
		token    prose.Token
		expected *BlockMapType
	}{
		{
			"Test Block UpdateBlockMap",
			&blockMapType1,
			miniStem,
			token,
			&blockMapType2,
		},
	}
	for _, test := range testBlockMapType {
		if got := test.receiver.UpdateBlockMap(test.miniStem, test.token); got != nil || !reflect.DeepEqual(test.receiver, test.expected) {
			t.Errorf("Failed! %s : \nerr: %v \nOR %v does not match expected : \n%v\n", test.tname, got, test.receiver, test.expected)
		}
	}

}
