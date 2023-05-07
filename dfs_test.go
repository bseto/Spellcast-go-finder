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
		Letter("a"), Letter("a"), Letter("e"), Letter("g"), Letter("n"),
	},
	{
		Letter("g"), Letter("l"), Letter("u"), Letter("n"), Letter("v"),
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
			want: []Node{{"o", 1}, {"e", 3}, {"a", 6}, {"e", 7}, {"d", 8}},
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
			//os.WriteFile(tt.wantAdjacencyFromJSON, b, 777)

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
