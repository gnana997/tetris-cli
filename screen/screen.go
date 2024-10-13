package screen

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

var colors = []termbox.Attribute{
	termbox.ColorBlack,
	termbox.ColorBlue,
	termbox.ColorCyan,
	termbox.ColorWhite,
	termbox.ColorYellow,
	termbox.ColorMagenta,
	termbox.ColorLightGray,
	termbox.ColorRed,
}

type gameScreen struct {
}

func (g *gameScreen) Render(board [][]int) {
	err := termbox.Clear(termbox.ColorGreen, termbox.ColorGreen)
	if err != nil {
		return
	}

	offsetY := 4
	offsetX := 4
	cellWidth := 2

	for y, row := range board {
		for x, col := range row {
			color := colors[col]
			for i := 0; i < cellWidth; i++ {
				termbox.SetCell(x*cellWidth+offsetX+i, y+offsetY, ' ', color, color)
			}
		}
	}

	err = termbox.Flush()
	if err != nil {
		return
	}
}

func (g *gameScreen) RenderAscii(board [][]int) {
	fmt.Println("\n=======")
	for _, row := range board {
		for _, col := range row {
			if col > 0 {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func New() *gameScreen {
	return &gameScreen{}
}
