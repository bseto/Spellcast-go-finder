package main

import (
	"reflect"
	"testing"
)

func TestNewBoard(t *testing.T) {
	type args struct {
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want *Board
	}{
		{
			name: "hello",
			args: args{
				width:  5,
				height: 5,
			},
			want: &Board{
				width:  5,
				height: 5,
				Tiles: [][]BoardTile{
					{
						BoardTile{Letter: "_"}, BoardTile{Letter: "_"}, BoardTile{Letter: "_"}, BoardTile{Letter: "_"}, BoardTile{Letter: "_"},
					},
					{
						BoardTile{Letter: "_"}, BoardTile{Letter: "_"}, BoardTile{Letter: "_"}, BoardTile{Letter: "_"}, BoardTile{Letter: "_"},
					},
					{
						BoardTile{Letter: "_"}, BoardTile{Letter: "_"}, BoardTile{Letter: "_"}, BoardTile{Letter: "_"}, BoardTile{Letter: "_"},
					},
					{
						BoardTile{Letter: "_"}, BoardTile{Letter: "_"}, BoardTile{Letter: "_"}, BoardTile{Letter: "_"}, BoardTile{Letter: "_"},
					},
					{
						BoardTile{Letter: "_"}, BoardTile{Letter: "_"}, BoardTile{Letter: "_"}, BoardTile{Letter: "_"}, BoardTile{Letter: "_"},
					},
				},
				VisualTiles: [][]string{
					{"_", "_", "_", "_", "_"},
					{"_", "_", "_", "_", "_"},
					{"_", "_", "_", "_", "_"},
					{"_", "_", "_", "_", "_"},
					{"_", "_", "_", "_", "_"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoard(tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
