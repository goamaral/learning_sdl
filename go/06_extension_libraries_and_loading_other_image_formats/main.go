/*
	Extension Libraries and Loading Other Image Formats
	https://lazyfoo.net/tutorials/SDL/06_extension_libraries_and_loading_other_image_formats
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

	surface := srf.Load("../resources/images/png_loaded.png", window.GetSurface().Format)
	eng.SaveSurface("pngLoaded", &surface)

	window.RenderSurface("pngLoaded", false)
	eng.ProcessEvents(nil)
	sdl.Delay(2000)
}
