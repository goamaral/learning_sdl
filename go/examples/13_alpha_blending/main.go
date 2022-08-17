/*
	Alpha Blending
	https://lazyfoo.net/tutorials/SDL/13_alpha_blending/index.php
*/

package main

import (
	"github.com/Goamaral/learning_sdl/engine"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"github.com/veandco/go-sdl2/sdl"
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
	window, err := engine.CreateWindow(false)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create window")
		return
	}
	defer window.Destroy()

	// Load surfaces
	fadeInSurface, err := window.LoadSurface("../../../resources/images/fadein.png")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load fade in surface")
		return
	}
	fadeOutSurface, err := window.LoadSurface("../../../resources/images/fadeout.png")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load fade out surface")
		return
	}

	// Convert surfaces to textures
	fadeInTexture, err := window.ConvertSurfaceToTexture(fadeInSurface)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert fade in surface to texture")
		return
	}
	fadeOutTexture, err := window.ConvertSurfaceToTexture(fadeOutSurface)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert fade out surface to texture")
		return
	}

	// Set fade out texture blend mode
	err = fadeOutTexture.SetBlendMode(sdl.BLENDMODE_BLEND)
	if err != nil {
		log.Error().Err(err).Msg("Failed to set fade out texture blend mode")
		return
	}

	var alpha uint8 = 255

	engine.EventLoop(func(getEvent func() sdl.Event) bool {
		for sdlEvent := getEvent(); sdlEvent != nil; sdlEvent = getEvent() {
			switch eventType := sdlEvent.(type) {
			// Key pressed
			case *sdl.KeyboardEvent:
				switch eventType.Keysym.Sym {
				case sdl.K_w:
					alpha = uint8(lo.Min([]int16{255, int16(alpha) + 32}))
				case sdl.K_s:
					alpha = uint8(lo.Max([]int16{0, int16(alpha) - 32}))
				}
			}
		}

		// Reset renderer
		err = window.Renderer.Reset()
		if err != nil {
			log.Error().Err(err).Msg("Failed to reset renderer")
			return false
		}

		// Set fade out texture alpha
		err = fadeOutTexture.SetAlphaMod(alpha)
		if err != nil {
			log.Error().Err(err).Msg("Failed to set fade out texture alpha mod")
			return false
		}

		// Render
		err = window.Renderer.RenderTexture(engine.RenderContext{}, &fadeInTexture, 0, 0, nil)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render fade in texture")
			return false
		}
		err = window.Renderer.RenderTexture(engine.RenderContext{}, &fadeOutTexture, 0, 0, nil)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render fade out texture")
			return false
		}

		window.Renderer.Present()

		return true
	})
}
