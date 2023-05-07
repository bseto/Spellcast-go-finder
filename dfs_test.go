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
		"I", "O", "B", "E", "T",
	},
	{
		"E", "A", "E", "D", "C",
	},
	{
		"E", "V", "R", "D", "D",
	},
	{
		"A", "A", "E", "G", "N",
	},
	{
		"G", "L", "U", "N", "V",
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
			want: []Node{{"O", 1}, {"E", 3}, {"A", 6}, {"E", 7}, {"D", 8}},
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
