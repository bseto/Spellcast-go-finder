package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

var exampleWordMatrix = [][]BoardTile{
	{
		Letter("i"), Letter("o"), Letter("b"), Letter("e"), Letter("t"),
	},
	{
		Letter("e"), Letter("a"), Letter("e"), Letter("d"), Letter("c"),
	},
	{
		Letter("e"), Letter("v"), Letter("r"), Letter("d"), Letter("d"),
	},
	{
		Letter("a"), Letter("a"), LetterMult("e", 2), Letter("g"), Letter("n"),
	},
	{
		Letter("g"), Letter("l"), Letter("u"), Letter("n"), LetterDouble("v"),
	},
}

var exampleWordMatrix2 = [][]BoardTile{
	{
		LetterMult("o", 3), Letter("d"), Letter("p"), Letter("u"), Letter("v"),
	},
	{
		Letter("n"), Letter("r"), Letter("b"), Letter("a"), Letter("g"),
	},
	{
		Letter("o"), Letter("w"), Letter("e"), Letter("e"), Letter("l"),
	},
	{
		Letter("o"), Letter("f"), Letter("s"), Letter("n"), Letter("e"),
	},
	{
		Letter("x"), Letter("u"), Letter("u"), Letter("e"), Letter("a"),
	},
}

func TestGetNeighbors(t *testing.T) {
	type args struct {
		wordMatrix [][]BoardTile
		row        int
		col        int
	}
	tests := []struct {
		name string
		args args
		want []Node
	}{
		{
			name: "Standard Board",
			args: args{
				wordMatrix: exampleWordMatrix,
				row:        0, col: 2, // B on the first row
			},
			want: []Node{
				{Letter: "o", AdjacencyAddress: 1, Multiplier: 1},
				{Letter: "e", AdjacencyAddress: 3, Multiplier: 1},
				{Letter: "a", AdjacencyAddress: 6, Multiplier: 1},
				{Letter: "e", AdjacencyAddress: 7, Multiplier: 1},
				{Letter: "d", AdjacencyAddress: 8, Multiplier: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNeighbors(tt.args.wordMatrix, tt.args.row, tt.args.col); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNeighbors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAdjacenyMatrix(t *testing.T) {
	type args struct {
		wordMatrix [][]BoardTile
	}
	tests := []struct {
		name                  string
		args                  args
		wantAdjacencyFromJSON string
	}{
		{
			name:                  "Standard Board",
			args:                  args{exampleWordMatrix},
			wantAdjacencyFromJSON: filepath.Join("testdata", "adjacency.golden.json"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//gotAdjacency := ToAdjacenyMatrix(tt.args.wordMatrix)
			//b, _ := json.Marshal(gotAdjacency)
			//os.WriteFile(tt.wantAdjacencyFromJSON, b, 0777)

			b, err := os.ReadFile(tt.wantAdjacencyFromJSON)
			if err != nil {
				t.Fatalf("unable to setup test: %v", err)
			}
			var expectedAdjacency [][]Node
			err = json.Unmarshal(b, &expectedAdjacency)
			if err != nil {
				t.Fatalf("unable to unmarshal to expected adjacency: %v", err)
			}
			if gotAdjacency := ToAdjacenyMatrix(tt.args.wordMatrix); !reflect.DeepEqual(gotAdjacency, expectedAdjacency) {
				t.Errorf("ToAdjacenyMatrix() = %v, want %v", gotAdjacency, expectedAdjacency)
			}
		})
	}
}
