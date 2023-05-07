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
	boardMatrix     [][]string
	adjacencyMatrix [][]Node
	words           []string
}

func NewSpellCastFinder(trie Trie) *SpellCastFinder {
	return &SpellCastFinder{
		trie: trie,
	}
}

func (s *SpellCastFinder) FindSolution() {

}

// DFS will return all the max length strings it can from the tile.
// Max length as in if there was a partial word zig, but then continued until
// it became zigzag, then only zigzag will be returned. if another branch of
// possibilities also include zigs, then zigs, and zigzag would be returned
func (s *SpellCastFinder) DFSRecursive(tile Node, currentWord string, visited map[int]bool) {
	if s.trie.Get(currentWord) {
		s.words = append(s.words, currentWord)
	}

	visited[tile.AdjacencyAddress] = true
	for _, letter := range s.adjacencyMatrix[tile.AdjacencyAddress] {
		potentialWord := currentWord + letter.Letter
		if !s.trie.Prefix(potentialWord) {
			continue
		}

		if !visited[letter.AdjacencyAddress] {
			copyOfVisited := make(map[int]bool)
			for k, v := range visited {
				copyOfVisited[k] = v
			}
			s.DFSRecursive(letter, potentialWord, copyOfVisited)
		}
	}
}
