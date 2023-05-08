package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var exampleWordMatrix3 = [][]BoardTile{
	{
		Letter("a"), Letter("n"), Letter("t"), LetterMult("u", 2), Letter("v"),
	},
	{
		LetterDouble("n"), Letter("w"), Letter("r"), Letter("d"), Letter("r"),
	},
	{
		Letter("t"), Letter("h"), Letter("i"), Letter("i"), Letter("l"),
	},
	{
		Letter("a"), Letter("a"), Letter("a"), Letter("p"), LetterDouble("q"),
	},
	{
		Letter("s"), Letter("n"), Letter("i"), Letter("e"), Letter("a"),
	},
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	board := NewBoard(5, 5)
	table := widgets.NewTable()
	table.Rows = board.GetVisualTiles()
	table.TextStyle = ui.NewStyle(ui.ColorWhite)
	table.TextAlignment = ui.AlignCenter
	table.RowSeparator = true
	table.BorderStyle = ui.NewStyle(ui.ColorWhite)
	table.SetRect(0, 0, 60, 10)
	table.FillRow = true
	table.RowStyles[0] = ui.NewStyle(ui.ColorWhite)
	table.RowStyles[2] = ui.NewStyle(ui.ColorWhite)
	table.RowStyles[3] = ui.NewStyle(ui.ColorWhite)
	table.RowStyles[4] = ui.NewStyle(ui.ColorWhite)
	table.RowStyles[5] = ui.NewStyle(ui.ColorWhite)
	ui.Render(table)
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "j", "<Down>":
		case "k", "<Up>":
		case "h", "<Left>":
			board.MoveLeft()
		case "l", "<Right>":
			board.MoveRight()
		case "q", "<C-c>":
			return
		}
		ui.Render(table)
	}

	//trie, err := NewTrie("words.txt")
	//if err != nil {
	//fmt.Printf("unable to open words.txt: %v", err)
	//return
	//}
	//finder := NewSpellCastFinder(trie, exampleWordMatrix3)
	//scores := finder.FindSolutionWithSwap()
	//scores = scores[len(scores)-5:]

	//a, err := json.Marshal(scores)
	//if err != nil {
	//fmt.Printf("unable to write to file: %v", err)
	//return
	//}
	//err = os.WriteFile("scores.json", a, 0777)
	//if err != nil {
	//fmt.Printf("unable to write to file: %v", err)
	//return
	//}
}
