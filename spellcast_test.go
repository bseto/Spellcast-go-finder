package main

import (
	"reflect"
	"testing"
)

func TestSpellCastFinder_DFS(t *testing.T) {
	type args struct {
		tile Node
	}
	tests := []struct {
		name string
		s    SpellCastFinder
		args args
		want [][]Node
	}{
		{
			name: "test with a tile that will generate a large word, 'devalued'",
			s: SpellCastFinder{
				trie: func(t *testing.T) Trie {
					trie, err := NewTrie("words.txt")
					if err != nil {
						t.Fatalf("unable to setup trie: %v", err)
					}
					return trie
				}(t),
				adjacencyMatrix: ToAdjacenyMatrix(exampleWordMatrix),
			},
			args: args{tile: Node{"D", 8}},
			want: [][]Node{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.DFSRecursive(tt.args.tile, "D", map[int]bool{})
			got := tt.s.words
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpellCastFinder.DFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
