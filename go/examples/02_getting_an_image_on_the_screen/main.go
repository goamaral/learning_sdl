/*
	Getting an Image on the Screen
	https://lazyfoo.net/tutorials/SDL/02_getting_an_image_on_the_screen
*/

package main

import (
	"github.com/rs/zerolog/log"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/Goamaral/learning_sdl/engine"
)

func main() {
	err := engine.Init()
	if err != nil {
		log.Error().Err(err).Msg("Failed to initialized engine")
		return
	}
	defer engine.Quit()

	window, err := engine.CreateWindow()
	if err != nil {
		log.Error().Err(err).Msg("Failed to create window")
		return
	}
	defer window.Destroy()

	surfaceId, err := window.LoadSurface("../../../resources/images/hello_world.bmp")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load surface")
		return
	}

	window.RenderSurface(surfaceId, false)
	engine.ProcessEvents(nil)
	sdl.Delay(2000)
}
