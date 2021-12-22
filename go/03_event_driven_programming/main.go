/*
	Event Driven Programming
	https://lazyfoo.net/tutorials/SDL/03_event_driven_programming
*/

package main

import (
	eng "github.com/Goamaral/learning_sdl/engine"
	win "github.com/Goamaral/learning_sdl/window"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	eng.Init()
	defer eng.Quit()

	window := win.Create()
	defer window.Destroy()

	eng.Loop(func(event sdl.Event) bool {
		switch eventType := event.(type) {
		// Key pressed
		case *sdl.KeyboardEvent:
			switch eventType.Keysym.Sym {
			case sdl.K_ESCAPE:
				return false
			}
		}
		return true
	})
}
