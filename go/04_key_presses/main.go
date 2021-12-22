/*
	Key Presses
	https://lazyfoo.net/tutorials/SDL/04_key_presses
*/

package main

import (
	eng "github.com/Goamaral/learning_sdl/engine"
	srf "github.com/Goamaral/learning_sdl/surface"
	win "github.com/Goamaral/learning_sdl/window"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	eng.Init()
	defer eng.Quit()

	window := win.Create()
	defer window.Destroy()

	surfacePaths := map[string]string{
		"default": "../resources/images/default.bmp",
		"up":      "../resources/images/up.bmp",
		"right":   "../resources/images/right.bmp",
		"down":    "../resources/images/down.bmp",
		"left":    "../resources/images/left.bmp",
	}

	for name, path := range surfacePaths {
		surface := srf.LoadFromBmp(path, window.GetSurface().Format)
		eng.SaveSurface(name, &surface)
	}

	window.RenderSurface("default")

	eng.Loop(func(event sdl.Event) bool {
		switch eventType := event.(type) {
		// Key pressed
		case *sdl.KeyboardEvent:
			switch eventType.Keysym.Sym {
			case sdl.K_UP:
				window.RenderSurface("up")
			case sdl.K_RIGHT:
				window.RenderSurface("right")
			case sdl.K_DOWN:
				window.RenderSurface("down")
			case sdl.K_LEFT:
				window.RenderSurface("left")
			default:
				window.RenderSurface("default")
			}
		}
		return true
	})
}
