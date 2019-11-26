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

func (node *TrieNode) Insert(word string) {
	for _, c := range []byte(word) {
		if node.next[c-'a'] == nil {
			node.next[c-'a'] = &TrieNode{}
		}
		node = node.next[c-'a']
	}
	node.word = word
}

func (node *TrieNode) Next(char byte) *TrieNode {
	if char >= 'a' && char <= 'z' {
		return node.next[char-'a']
	}
	return nil
}

type Puzzle struct {
	row, col int
	words    [][]byte
}

func (puzzle *Puzzle) Search(row, col int, node *TrieNode, dir Direction) *Result {
	if node == nil {
		return nil
	}
	if node.word != "" {
		return &Result{[2]int{row, col}, node.word}
	}
	if row, col = row+dir[0], col+dir[1]; (row >= 0 && row < puzzle.row) && (col >= 0 && col < puzzle.col) {
		return puzzle.Search(row, col, node.Next(puzzle.words[row][col]), dir)
	}
	return nil
}

func Solve(words []string, puzzleWords []string) (map[string][2][2]int, error) {
	trieRoot := &TrieNode{}
	for _, word := range words {
		trieRoot.Insert(word)
	}
	puzzle := &Puzzle{
		row: len(puzzleWords[0]),
		col: len(puzzleWords),
	}
	puzzle.words = make([][]byte, puzzle.row)
	for i := range puzzle.words {
		puzzle.words[i] = make([]byte, puzzle.col)
		for j := range puzzle.words[i] {
			puzzle.words[i][j] = puzzleWords[j][i]
		}
	}

	res := map[string][2][2]int{}
	for row := range puzzle.words {
		for col := range puzzle.words[row] {
			if node := trieRoot.Next(puzzle.words[row][col]); node != nil {
				for _, dir := range Directions {
					if result := puzzle.Search(row, col, node, dir); result != nil {
						res[result.word] = [2][2]int{{row, col}, result.endPos}
						if len(res) == len(words) {
							return res, nil
						}
					}
				}
			}
		}
	}

	return nil, ErrNoWordInPuzzle
}