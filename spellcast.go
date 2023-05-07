package main

type BoardTile struct {
	Letter      string
	Multiplier  int
	DoublePoint bool
}

func Letter(letter string) BoardTile {
	return BoardTile{Letter: letter}
}

func LetterMult(letter string, multiplier int) BoardTile {
	return BoardTile{Letter: letter, Multiplier: multiplier}
}

func LetterDouble(letter string) BoardTile {
	return BoardTile{Letter: letter, DoublePoint: true}
}

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

func NewSpellCastFinder(trie Trie, boardMatrix [][]BoardTile) *SpellCastFinder {
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

// DFS will return all words from a tile that are valid. The only optimization in this DFS is
// that it'll stop if a potential word is no longer a prefix in the Trie, as in the
// potential word is no longer potentially a valid word
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
