package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/derekparker/trie"
)

// TrieWrapper wraps the derekparker/trie implementation. I can maybe implement
// my own Trie later
type TrieWrapper struct {
	trie.Trie
}

func NewTrie(filepath string) (Trie, error) {
	wrapper := TrieWrapper{
		Trie: *trie.New(),
	}

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("unable to read file: %v\n", err)
		return &wrapper, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wrapper.Add(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error scanning file: %v\n", err)
		return &wrapper, err
	}

	return &wrapper, nil
}

func (t *TrieWrapper) Add(key string) {
	t.Trie.Add(key, 0)
}

func (t *TrieWrapper) Get(key string) bool {
	_, ok := t.Trie.Find(strings.ToLower(key))
	return ok
}
func (t *TrieWrapper) Delete(key string) {
	t.Trie.Remove(key)
}

// Will return true if there are possible paths with the current prefix
func (t *TrieWrapper) Prefix(key string) bool {
	return t.Trie.HasKeysWithPrefix(strings.ToLower(key))
}
