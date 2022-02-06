package profile

import (
	"sort"

	"github.com/jdkato/prose/v3"
)

/*
Block has the minified stem of words extracted
*/
type Block struct {
	MiniStem string     `json:"mini_stem"`
	Sources  []WordSeen `json:"sources"`
	Weight   int        `json:"weight"`
	Count    int        `json:"count"`
}

/*
BlockMapType is for Mapped Blocks to be used as method receivers
*/
type BlockMapType struct {
	BlockMap map[string]Block `json:"block_map,omitempty"`
}

/*
UpdateBlockMap updates the BlockMap using BlockMapType as receiver
*/
func (blockMapType *BlockMapType) UpdateBlockMap(miniStem string, token prose.Token) error {

	var blockMap map[string]Block = blockMapType.BlockMap
	var block Block
	var ok bool
	var wordSeen = WordSeen{Word: token.Text, Seen: 1}
	var wordsSeen []WordSeen = []WordSeen{wordSeen}

	block, ok = blockMap[miniStem]
	if !ok {
		block = Block{MiniStem: miniStem, Weight: 1, Count: 1}

		block.Sources = wordsSeen
		blockMap[miniStem] = block
		blockMapType.BlockMap = blockMap
		return nil
	}
	var hasWordAt int = -1
	for i, v := range block.Sources {
		if wordSeen.Word == v.Word {
			hasWordAt = i
			wordSeen.Seen = v.Seen + 1
		}
	}
	if hasWordAt == -1 {
		block.Sources = append(block.Sources, wordSeen)
	} else {
		block.Sources[hasWordAt] = wordSeen
	}
	block.Count = block.Count + 1
	blockMap[miniStem] = block

	blockMapType.BlockMap = blockMap

	return nil
}

/*
MapKeysAsStrings returns slice of string keys from BlockMapType.BlockMap
*/
func (blockMapType *BlockMapType) MapKeysAsStrings() []string {
	blockMap := blockMapType.BlockMap
	keys := make([]string, len(blockMap))
	sort.Strings(keys)
	i := 0
	for k := range blockMap {
		keys[i] = k
		i++
	}
	return keys
}

/*
Combine combines blockMaps
*/
func (blockMapType *BlockMapType) Combine(inBlockMapTypes []BlockMapType) error {
	// Get keys for *blockMapType

	baseBlockMap := blockMapType.BlockMap

	for _, inBlockMapType := range inBlockMapTypes {
		inBlockMap := inBlockMapType.BlockMap
		inBlockMapKeys := inBlockMapType.MapKeysAsStrings()
		for _, inBlockKey := range inBlockMapKeys {
			inBlock, _ := inBlockMap[inBlockKey]
			baseBlock, ok := baseBlockMap[inBlockKey]
			if ok {
				baseBlock.Count++
				// update baseBlock.Sources
				baseBlockSources := WordsSeenType{WordsSeen: baseBlock.Sources}
				inBlockSource := WordsSeenType{WordsSeen: inBlock.Sources}
				baseBlockSources.Combine(inBlockSource)
				baseBlock.Sources = baseBlockSources.WordsSeen
				baseBlockMap[inBlockKey] = baseBlock
			} else {
				baseBlockMap[inBlockKey] = inBlock
			}
		}
	}
	blockMapType.BlockMap = baseBlockMap
	return nil
}
