/*
	Optimized Surface Loading and Soft Stretching
	https://lazyfoo.net/tutorials/SDL/05_optimized_surface_loading_and_soft_stretching
*/

package main

import (
	"time"

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
	window, err := engine.NewWindow(640, 480, false)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create window")
		return
	}

	// Load surface
	surface, err := window.LoadSurface("../../../resources/images/stretch.bmp")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load surface")
		return
	}

	// Event loop
	engine.EventLoop(func(getEvent func() sdl.Event) bool {
		// Render surface not scaled
		err = window.RenderSurface(surface, false)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render not scaled surface")
			return false
		}
		time.Sleep(2 * time.Second)

		// Render surface scaled
		err = window.RenderSurface(surface, true)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render scaled surface")
			return false
		}
		engine.EventLoop(nil)
		time.Sleep(2 * time.Second)

		return false
	})
}
