package window

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	*sdl.Window
}

// Create window
func Create() Window {
	sdlWindow, err := sdl.CreateWindow("SDL Tutorial", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 480, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	window := Window{Window: sdlWindow}

	return window
}
