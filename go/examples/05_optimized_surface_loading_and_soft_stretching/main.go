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
	// Init engine
	err := engine.Init()
	if err != nil {
		log.Error().Err(err).Msg("Failed to initialized engine")
		return
	}
	defer engine.Quit()

	// Create window
	window, err := engine.CreateWindow()
	if err != nil {
		log.Error().Err(err).Msg("Failed to create window")
		return
	}
	defer window.Destroy()

	// Load surface
	surfaceId, err := window.LoadSurface("../../../resources/images/stretch.bmp")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load surface")
		return
	}

	// Render surface not scaled
	err = window.RenderSurface(surfaceId, false)
	if err != nil {
		log.Error().Err(err).Msg("Failed to render not scaled surface")
		return
	}
	engine.ProcessEvents(nil)
	sdl.Delay(2000)

	// Render surface scaled
	err = window.RenderSurface(surfaceId, true)
	if err != nil {
		log.Error().Err(err).Msg("Failed to render scaled surface")
		return
	}
	engine.ProcessEvents(nil)
	sdl.Delay(2000)
}
