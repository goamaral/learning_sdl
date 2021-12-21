/*
	Event Driven Programming
	https://lazyfoo.net/tutorials/SDL/03_event_driven_programming
*/

package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// Init SDL
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// Create window
	window, err := sdl.CreateWindow("SDL Tutorial", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 480, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// Main loop
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch eventType := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				switch eventType.Keysym.Sym {
				case sdl.K_ESCAPE:
					running = false
				}
			}
		}
	}
}
