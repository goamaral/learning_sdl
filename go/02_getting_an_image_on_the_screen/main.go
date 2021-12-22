/*
	Getting an Image on the Screen
	https://lazyfoo.net/tutorials/SDL/02_getting_an_image_on_the_screen
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

	surface := srf.LoadFromBmp("../resources/images/hello_world.bmp", window.GetSurface().Format)
	eng.SaveSurface("helloWorld", &surface)

	window.RenderRawSurface(&surface)

	eng.Loop(func(event sdl.Event) bool {
		return true
	})
}
