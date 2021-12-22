package window

import (
	srf "github.com/Goamaral/learning_sdl/surface"
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

// Get window surface
func (w *Window) GetSurface() srf.Surface {
	sdlSurface, err := w.Window.GetSurface()
	if err != nil {
		panic(err)
	}
	return srf.Surface{Surface: sdlSurface}
}

// Render surface
func (w *Window) RenderSurface(surface *srf.Surface) {
	err := surface.Blit(nil, w.GetSurface().Surface, nil)
	if err != nil {
		panic(err)
	}

	w.UpdateSurface()
}
