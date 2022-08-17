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

	// Load surfaces and map ids
	surfacePaths := map[string]string{
		"default": "../../../resources/images/default.bmp",
		"up":      "../../../resources/images/up.bmp",
		"right":   "../../../resources/images/right.bmp",
		"down":    "../../../resources/images/down.bmp",
		"left":    "../../../resources/images/left.bmp",
	}
	surfaceIds := map[string]uint32{
		"default": 0,
		"up":      0,
		"right":   0,
		"down":    0,
		"left":    0,
	}

	for name, path := range surfacePaths {
		surface, err := window.LoadSurface(path)
		surfaceIds[name] = surface.ID
		if err != nil {
			log.Error().Err(err).Msgf("Failed to load surface %s at %s", name, path)
			return
		}
	}

	// Render default surface
	err = window.RenderSurface(surfaceIds["default"], false)
	if err != nil {
		log.Error().Err(err).Msg("Failed to render default surface")
		return
	}

	// Run event loop
	engine.EventLoop(func(getEvent func() sdl.Event) bool {
		var surfaceIdToRender uint32

		// Events
		for sdlEvent := getEvent(); sdlEvent != nil; sdlEvent = getEvent() {
			switch event := sdlEvent.(type) {
			// Key press
			case *sdl.KeyboardEvent:
				switch event.Keysym.Sym {
				case sdl.K_UP:
					surfaceIdToRender = surfaceIds["up"]
				case sdl.K_RIGHT:
					surfaceIdToRender = surfaceIds["right"]
				case sdl.K_DOWN:
					surfaceIdToRender = surfaceIds["down"]
				case sdl.K_LEFT:
					surfaceIdToRender = surfaceIds["left"]
				default:
					surfaceIdToRender = surfaceIds["default"]
				}
			}
		}

		// Render
		if surfaceIdToRender != 0 {
			err = window.RenderSurface(surfaceIdToRender, false)
			if err != nil {
				log.Error().Err(err).Msg("Failed to render surface")
			}
		}

		return true
	})
}
