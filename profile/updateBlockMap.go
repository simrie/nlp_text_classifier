package profile

import (
	"github.com/jdkato/prose/v3"
)

/*
UpdateBlockMap updates the BlockMap using BlockMapType as receiver
*/
func (blockMapType *BlockMapType) UpdateBlockMap(miniStem string, token prose.Token) error {

	var blockMap map[string]Block = blockMapType.BlockMap
	var block Block
	var ok bool
	var wordSeen = WordSeen{Word: token.Text, Seen: 1}

	block, ok = blockMap[miniStem]
	if !ok {
		block = Block{MiniStem: miniStem, Weight: 1, Count: 1}
		block.Source = []WordSeen{wordSeen}
		blockMap[miniStem] = block
		return nil
	}
	var hasWordAt int = -1
	for i, v := range block.Source {
		if wordSeen.Word == v.Word {
			hasWordAt = i
			wordSeen.Seen = v.Seen + 1
		}
	}
	if hasWordAt == -1 {
		block.Source = append(block.Source, wordSeen)
	} else {
		block.Source[hasWordAt] = wordSeen
	}
	block.Count = block.Count + 1
	blockMap[miniStem] = block

	return nil
}
