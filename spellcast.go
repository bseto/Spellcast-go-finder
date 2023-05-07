package main

type Node struct {
	Letter           string
	AdjacencyAddress int
}

type Trie interface {
	Add(string)
	Get(string) bool
	Prefix(string) bool
	Delete(string)
}

type SpellCastFinder struct {
	trie            Trie
	adjacencyMatrix [][]Node
	words           []string
}

func NewSpellCastFinder(trie Trie, boardMatrix [][]string) *SpellCastFinder {
	return &SpellCastFinder{
		trie:            trie,
		adjacencyMatrix: ToAdjacenyMatrix(boardMatrix),
	}
}

func (s *SpellCastFinder) FindSolution() []Score {
	numOfTiles := len(s.adjacencyMatrix)
	for i := 0; i < numOfTiles; i++ {
		s.DFSRecursive(i, "", map[int]bool{})
	}
	s.words = RemoveDuplicates(s.words)
	return CalculateAndSortByScore(s.words)
}

// DFS will return all the max length strings it can from the tile.
// Max length as in if there was a partial word zig, but then continued until
// it became zigzag, then only zigzag will be returned. if another branch of
// possibilities also include zigs, then zigs, and zigzag would be returned
func (s *SpellCastFinder) DFSRecursive(tileNum int, currentWord string, visited map[int]bool) {
	if s.trie.Get(currentWord) {
		s.words = append(s.words, currentWord)
	}

	visited[tileNum] = true
	for _, letter := range s.adjacencyMatrix[tileNum] {
		potentialWord := currentWord + letter.Letter
		if !s.trie.Prefix(potentialWord) {
			continue
		}

		if !visited[letter.AdjacencyAddress] {
			copyOfVisited := make(map[int]bool)
			for k, v := range visited {
				copyOfVisited[k] = v
			}
			s.DFSRecursive(letter.AdjacencyAddress, potentialWord, copyOfVisited)
		}
	}
}
