package main

import "strings"

// ToAdjacenyMatrix will create a tileNumCursor which represents which tile we
// are on. Starting from the top left and traveling through the row to the right
// until we hit the next col. Then we start at the very left of the next col.
func ToAdjacenyMatrix(wordMatrix [][]BoardTile) (adjacency [][]Node) {
	tileNumCursor := 0
	width, height := GetWidthAndHeight(wordMatrix)
	adjacency = make([][]Node, width*height)
	for i, row := range wordMatrix {
		for j := range row {
			adjacency[tileNumCursor] = GetNeighbors(wordMatrix, i, j)
			tileNumCursor++
		}
	}
	return
}

func GetNeighbors(wordMatrix [][]BoardTile, row, col int) []Node {
	width, height := GetWidthAndHeight(wordMatrix)
	var Nodes []Node
	for rowDeviation := -1; rowDeviation < 2; rowDeviation++ {
		if rowDeviation == -1 && row == 0 {
			continue // skip the left side of the node
		}
		if rowDeviation == 1 && row == (width-1) {
			continue // skip the right side of the node
		}
		for colDeviation := -1; colDeviation < 2; colDeviation++ {
			if colDeviation == -1 && col == 0 {
				continue // skip above the node
			}
			if colDeviation == 1 && col == (height-1) {
				continue // skip below of the node
			}
			if colDeviation == 0 && rowDeviation == 0 {
				continue // skip itself
			}
			Nodes = append(
				Nodes,
				Node{
					Letter:           wordMatrix[row+rowDeviation][col+colDeviation].Letter,
					AdjacencyAddress: (height * (row + rowDeviation)) + col + colDeviation,
				},
			)
		}
	}
	return Nodes
}

// GetWidthAndHeight will assume the following format. outerslice represents rows, inside represents col
// [[a, b, c],
// [d, e, f],
// [g, h, i]]
// GetWidthAndHeight will return the width and height of the word Matrix
func GetWidthAndHeight[T any](wordMatrix [][]T) (int, int) {
	height := len(wordMatrix)
	width := len(wordMatrix[0]) // honestly, if there's no width, just panic. I'm not gonna return an error here
	return width, height
}

func StringFromSlice(letters []string) string {
	var b strings.Builder
	for _, letter := range letters {
		b.WriteString(letter)
	}
	return b.String()
}

func RemoveDuplicates(slice []string) []string {
	tempMap := make(map[string]bool)
	newList := []string{}
	for _, s := range slice {
		if _, ok := tempMap[s]; !ok {
			tempMap[s] = true
			newList = append(newList, s)
		}
	}
	return newList
}
