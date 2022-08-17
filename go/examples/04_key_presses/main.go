/*
	Key Presses
	https://lazyfoo.net/tutorials/SDL/04_key_presses
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

	// Load surfaces
	surfacePaths := map[string]string{
		"default": "../../../resources/images/default.bmp",
		"up":      "../../../resources/images/up.bmp",
		"right":   "../../../resources/images/right.bmp",
		"down":    "../../../resources/images/down.bmp",
		"left":    "../../../resources/images/left.bmp",
	}
	surfaces := map[string]*engine.Surface{}
	for name, path := range surfacePaths {
		surface, err := window.LoadSurface(path)
		surfaces[name] = surface
		if err != nil {
			log.Error().Err(err).Msgf("Failed to load surface %s at %s", name, path)
			return
		}
	}

	// Render default surface
	err = window.RenderSurface(surfaces["default"], false)
	if err != nil {
		log.Error().Err(err).Msg("Failed to render default surface")
		return
	}

	// Run event loop
	engine.EventLoop(func(getEvent func() sdl.Event) bool {
		var surface *engine.Surface

		// Events
		for sdlEvent := getEvent(); sdlEvent != nil; sdlEvent = getEvent() {
			switch event := sdlEvent.(type) {
			// Key press
			case *sdl.KeyboardEvent:
				switch event.Keysym.Sym {
				case sdl.K_UP:
					surface = surfaces["up"]
				case sdl.K_RIGHT:
					surface = surfaces["right"]
				case sdl.K_DOWN:
					surface = surfaces["down"]
				case sdl.K_LEFT:
					surface = surfaces["left"]
				default:
					surface = surfaces["default"]
				}
			}
		}

		// Render
		if surface != nil {
			err = window.RenderSurface(surface, false)
			if err != nil {
				log.Error().Err(err).Msg("Failed to render surface")
			}
		}

		return true
	})
}
