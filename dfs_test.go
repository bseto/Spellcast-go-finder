package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

var exampleWordMatrix = [][]string{
	{
		"i", "o", "b", "e", "t",
	},
	{
		"e", "a", "e", "d", "c",
	},
	{
		"e", "v", "r", "d", "d",
	},
	{
		"a", "a", "e", "g", "n",
	},
	{
		"g", "l", "u", "n", "v",
	},
}

var examplewordmatrix1 = [][]string{
	{
		"z", "f", "t", "f", "b",
	},
	{
		"n", "i", "g", "l", "x",
	},
	{
		"t", "r", "i", "u", "j",
	},
	{
		"e", "v", "o", "a", "y",
	},
	{
		"s", "r", "e", "r", "o",
	},
}

func TestGetNeighbors(t *testing.T) {
	type args struct {
		wordMatrix [][]string
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
		wordMatrix [][]string
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
