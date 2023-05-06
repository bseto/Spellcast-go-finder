package main

import (
	"reflect"
	"testing"
)

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
				wordMatrix: [][]string{
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
				},
				row: 0, col: 2, // B on the first row
			},
			want: []Node{Node{Letter: "O"}, Node{Letter: "E"}, Node{Letter: "A"}, Node{Letter: "E"}, Node{Letter: "D"}},
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
