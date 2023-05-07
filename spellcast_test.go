package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestSpellCastFinder_FindSolution(t *testing.T) {
	tests := []struct {
		name       string
		s          *SpellCastFinder
		goldenFile string
	}{
		{
			name: "full test with example word matrix",
			s: func(t *testing.T) *SpellCastFinder {
				trie, err := NewTrie("words.txt")
				if err != nil {
					t.Fatalf("unable to setup trie: %v", err)
				}
				return NewSpellCastFinder(trie, exampleWordMatrix)
			}(t),
			goldenFile: filepath.Join("testdata", "examplewordmatrix.SpellCastFinder.json"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := os.ReadFile(tt.goldenFile)
			if err != nil {
				t.Fatalf("unable to setup test: %v", err)
			}
			var expectedWordList []string
			err = json.Unmarshal(b, &expectedWordList)
			if err != nil {
				t.Fatalf("unable to unmarshal to expected adjacency: %v", err)
			}

			tt.s.FindSolution()

			if !reflect.DeepEqual(tt.s.words, expectedWordList) {
				t.Errorf("ToAdjacenyMatrix() = %v, want %v", tt.s.words, expectedWordList)
			}

			//b, _ := json.Marshal(tt.s.words)
			//os.WriteFile(tt.goldenFile, b, 777)
		})
	}
}
