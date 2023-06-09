package main

import (
	"strings"
)

type Node struct {
	Letter           string
	AdjacencyAddress int
	Multiplier       int
	IsDoublePoint    bool
}

type Trie interface {
	Add(string)
	Get(string) bool
	Prefix(string) bool
	Delete(string)
}

type SpellCastFinder struct {
	trie            Trie
	boardMatrix     [][]BoardTile
	adjacencyMatrix [][]Node
	words           []NodeWord
	debugger        []string
}

type NodeWord struct {
	Word                            Nodes
	AdjacencyAddressToSwappedLetter map[int]string // if there is any swapping happening, report it here
}

func (n NodeWord) Append(node Node) NodeWord {
	copyOfExisting := make(Nodes, len(n.Word))
	copy(copyOfExisting, n.Word)
	return NodeWord{
		Word: append(copyOfExisting, node),
	}
}

type Nodes []Node

func (n Nodes) ToString() string {
	var builder strings.Builder
	for _, node := range n {
		builder.WriteString(node.Letter)
	}
	return builder.String()
}

func NewSpellCastFinder(trie Trie, boardMatrix [][]BoardTile) *SpellCastFinder {
	return &SpellCastFinder{
		trie:        trie,
		boardMatrix: boardMatrix,
	}
}

func (s *SpellCastFinder) FindSolution() []Score {
	s.adjacencyMatrix = ToAdjacenyMatrix(s.boardMatrix)
	numOfTiles := len(s.adjacencyMatrix)
	for i := 0; i < numOfTiles; i++ {
		s.DFSRecursive(i, NodeWord{}, map[int]bool{})
	}
	s.words = RemoveDuplicates(s.words)
	return CalculateAndSortByScore(s.words)
}

func (s *SpellCastFinder) FindSolutionWithSwap() []Score {
	width, height := GetWidthAndHeight(s.boardMatrix)
	// for every tile, we want to swap the letter
	for col := 0; col < height; col++ {
		for row := 0; row < width; row++ {
			for letter := 'a'; letter <= 'z'; letter++ {
				// bug, need to reset the board
				oldLetter := s.boardMatrix[col][row]
				s.boardMatrix[col][row].Letter = string(letter)
				s.adjacencyMatrix = ToAdjacenyMatrix(s.boardMatrix)
				numOfTiles := len(s.adjacencyMatrix)
				for i := 0; i < numOfTiles; i++ {
					s.DFSRecursive(i, NodeWord{}, map[int]bool{})
				}
				s.boardMatrix[col][row] = oldLetter
			}
		}
	}
	s.words = RemoveDuplicates(s.words)
	return CalculateAndSortByScore(s.words)
}

// DFS will return all words from a tile that are valid. The only optimization in this DFS is
// that it'll stop if a potential word is no longer a prefix in the Trie, as in the
// potential word is no longer potentially a valid word
func (s *SpellCastFinder) DFSRecursive(tileNum int, currentWord NodeWord, visited map[int]bool) {
	if s.trie.Get(currentWord.Word.ToString()) {
		s.words = append(s.words, currentWord)
		s.debugger = append(s.debugger, currentWord.Word.ToString())
	}

	visited[tileNum] = true
	for _, letter := range s.adjacencyMatrix[tileNum] {
		if !visited[letter.AdjacencyAddress] {
			potentialWord := currentWord.Append(letter)
			if !s.trie.Prefix(potentialWord.Word.ToString()) {
				continue
			}
			copyOfVisited := make(map[int]bool)
			for k, v := range visited {
				copyOfVisited[k] = v
			}
			s.DFSRecursive(letter.AdjacencyAddress, potentialWord, copyOfVisited)
		}
	}
}
