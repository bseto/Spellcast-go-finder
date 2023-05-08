package main

type BoardTile struct {
	Letter        string
	Multiplier    int
	IsDoublePoint bool
}

func Letter(letter string) BoardTile {
	return BoardTile{Letter: letter, Multiplier: 1}
}

func LetterMult(letter string, multiplier int) BoardTile {
	return BoardTile{Letter: letter, Multiplier: multiplier}
}

func LetterDouble(letter string) BoardTile {
	return BoardTile{Letter: letter, Multiplier: 1, IsDoublePoint: true}
}

type Board struct {
	// Board is Tiles[0] == first row
	// Tiles[0][5] == first row, 5th col
	Tiles       [][]BoardTile
	VisualTiles [][]string

	width  int
	height int

	cursorX int
	cursorY int
}

func NewBoard(width, height int) *Board {
	visual := make([][]string, height)
	board := make([][]BoardTile, height)
	for h := 0; h < height; h++ {
		visualRow := make([]string, width)
		row := make([]BoardTile, width)
		for w := 0; w < width; w++ {
			visualRow[w] = "_"
			row[w].Letter = "_"
		}
		visual[h] = visualRow
		board[h] = row
	}

	return &Board{
		width:       width,
		height:      height,
		Tiles:       board,
		VisualTiles: visual,
	}
}

func (b *Board) GetVisualTiles() [][]string {
	return b.VisualTiles
}

func (b *Board) MoveLeft() {
	if b.cursorX > 0 {
		b.VisualTiles[b.cursorY][b.cursorX] = b.Tiles[b.cursorY][b.cursorX].Letter
		b.cursorX--
		b.VisualTiles[b.cursorY][b.cursorX] = "X"
	}
}

func (b *Board) MoveRight() {
	if b.cursorX < b.width-1 {
		b.VisualTiles[b.cursorY][b.cursorX] = b.Tiles[b.cursorY][b.cursorX].Letter
		b.cursorX++
		b.VisualTiles[b.cursorY][b.cursorX] = "X"
	}
}
