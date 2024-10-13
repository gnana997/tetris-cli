package main

import (
	"flag"
	"github.com/nsf/termbox-go"
	"tetris-game/screen"
	"tetris-game/tetris"
	"time"
)

func main() {

	bdHeight := flag.Int("height", 20, "Board height")
	bdWidth := flag.Int("width", 10, "Board width")
	flag.Parse() // Create a new game

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()
	animationSpeed := time.Duration(50 * time.Millisecond)

	ticker := time.NewTimer(animationSpeed)

	// Create a new game
	game := tetris.NewGame(*bdHeight, *bdWidth)
	// Create a new screen
	scr := screen.New()

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Key == termbox.KeyArrowUp:
					game.Rotate()
				case ev.Key == termbox.KeyArrowDown:
					game.SpeedUp()
				case ev.Key == termbox.KeyArrowLeft:
					game.MoveLeft()
				case ev.Key == termbox.KeyArrowRight:
					game.MoveRight()
				case ev.Key == termbox.KeySpace:
					game.Fall()
				case ev.Ch == 'q':
					//quit
					return
				case ev.Ch == 's':
					//start
					game.Start()
				}
			}
		case <-ticker.C:
			scr.Render(game.GetBoard())
			ticker.Reset(animationSpeed)
		case <-game.FallSpeed.C:
			game.GameLoop()
		}
	}
}
