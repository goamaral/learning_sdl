/*
	Optimized Surface Loading and Soft Stretching
	https://lazyfoo.net/tutorials/SDL/05_optimized_surface_loading_and_soft_stretching
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

	surface := srf.Load("../resources/images/stretch.bmp", window.GetSurface().Format)
	eng.SaveSurface("stretch", &surface)

	window.RenderSurface("stretch", false)
	eng.ProcessEvents(nil)
	sdl.Delay(2000)

	window.RenderSurface("stretch", true)
	eng.ProcessEvents(nil)
	sdl.Delay(2000)
}
