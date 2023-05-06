package main

type Trie interface {
	Add(string)
	Get(string) bool
	Delete(string)
}

type BoardDFS interface {
	// Continue means to continue the DFS search
	StartNewTile() string
	// Break will back out of the branch and continue the DFS
	Break() string
}

type SpellCastFinder struct {
	trie  Trie
	board BoardDFS
}
