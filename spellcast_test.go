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
			scores := tt.s.FindSolution()
			//a, err := json.Marshal(scores)
			//if err != nil {
			//t.Fatalf("unable to write to file: %v", err)
			//}
			//err = os.WriteFile(tt.goldenFile, a, 0777)
			//if err != nil {
			//t.Fatalf("unable to write to file: %v", err)
			//}

			b, err := os.ReadFile(tt.goldenFile)
			if err != nil {
				t.Fatalf("unable to setup test: %v", err)
			}
			var expectedScores []Score
			err = json.Unmarshal(b, &expectedScores)
			if err != nil {
				t.Fatalf("unable to unmarshal to expected adjacency: %v", err)
			}

			if !reflect.DeepEqual(scores, expectedScores) {
				t.Errorf("ToAdjacenyMatrix() = %v, want %v", tt.s.words, expectedScores)
			}

		})
	}
}

func TestSpellCastFinder_ManualRun(t *testing.T) {
	tests := []struct {
		name string
		s    *SpellCastFinder
	}{
		{
			name: "Manual Run",
			s: func(t *testing.T) *SpellCastFinder {
				trie, err := NewTrie("words.txt")
				if err != nil {
					t.Fatalf("unable to setup trie: %v", err)
				}
				return NewSpellCastFinder(trie, exampleWordMatrix2)
			}(t),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//scores := tt.s.FindSolution()
			scores := tt.s.FindSolutionWithSwap()
			t.Logf("scores: %v", scores)
		})
	}
}
