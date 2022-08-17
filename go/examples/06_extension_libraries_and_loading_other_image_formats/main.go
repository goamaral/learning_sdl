/*
	Extension Libraries and Loading Other Image Formats
	https://lazyfoo.net/tutorials/SDL/06_extension_libraries_and_loading_other_image_formats
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
	window, err := engine.CreateWindow()
	if err != nil {
		log.Error().Err(err).Msg("Failed to create window")
		return
	}
	defer window.Destroy()

	// Load surface
	surface, err := window.LoadSurface("../../../resources/images/png_loaded.png")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load surface")
		return
	}

	// Process events
	engine.EventLoop(func(getEvent func() sdl.Event) bool {
		// Render surface
		err = window.RenderSurface(surface.ID, false)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render surface")
			return false
		}
		time.Sleep(2 * time.Second)

		return false
	})
}
