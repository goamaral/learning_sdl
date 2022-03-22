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

	surfacePaths := map[string]string{
		"default": "../../../resources/images/default.bmp",
		"up":      "../../../resources/images/up.bmp",
		"right":   "../../../resources/images/right.bmp",
		"down":    "../../../resources/images/down.bmp",
		"left":    "../../../resources/images/left.bmp",
	}
	surfaceIds := map[string]uint64{
		"default": 0,
		"up":      0,
		"right":   0,
		"down":    0,
		"left":    0,
	}

	for name, path := range surfacePaths {
		surfaceId, err := window.LoadSurface(path)
		surfaceIds[name] = surfaceId
		if err != nil {
			log.Error().Err(err).Msgf("Failed to load surface %s at %s", name, path)
			return
		}
	}

	window.RenderSurface(surfaceIds["default"], false)

	engine.EventLoop(func(event sdl.Event) bool {
		switch eventType := event.(type) {
		// Key pressed
		case *sdl.KeyboardEvent:
			switch eventType.Keysym.Sym {
			case sdl.K_UP:
				window.RenderSurface(surfaceIds["up"], false)
			case sdl.K_RIGHT:
				window.RenderSurface(surfaceIds["right"], false)
			case sdl.K_DOWN:
				window.RenderSurface(surfaceIds["down"], false)
			case sdl.K_LEFT:
				window.RenderSurface(surfaceIds["left"], false)
			default:
				window.RenderSurface(surfaceIds["default"], false)
			}
		}

		return true
	})
}
