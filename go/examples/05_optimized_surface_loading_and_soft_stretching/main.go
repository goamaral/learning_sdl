/*
	Optimized Surface Loading and Soft Stretching
	https://lazyfoo.net/tutorials/SDL/05_optimized_surface_loading_and_soft_stretching
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

	surfaceId, err := window.LoadSurface("../../../resources/images/stretch.bmp")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load surface")
		return
	}

	window.RenderSurface(surfaceId, false)
	engine.ProcessEvents(nil)
	sdl.Delay(2000)

	window.RenderSurface(surfaceId, true)
	engine.ProcessEvents(nil)
	sdl.Delay(2000)
}
