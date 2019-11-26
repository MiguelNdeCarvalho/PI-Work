package wordsearch

import (
	"errors"
)

var puzzleWords [][]byte

var (
	ErrNoWordInPuzzle = errors.New("no word in puzzle")
)

type Direction [2]int // [2]int{row_offset, col_offset}

var Directions = []Direction{
	{-1, 0},  // North
	{-1, 1},  // NorthEast
	{0, 1},   // East
	{1, 1},   // SouthEast
	{1, 0},   // South
	{1, -1},  // SouthWest
	{0, -1},  // West
	{-1, -1}, // NorthWest
}

type Result struct {
	endPos [2]int
	word   string
}

type TrieNode struct {
	word string
	next [26]*TrieNode
}
