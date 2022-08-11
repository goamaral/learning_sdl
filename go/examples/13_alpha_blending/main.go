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
	window, err := engine.CreateWindow()
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
	fadeInTexture, err := window.ConvertSurfaceToTexture(fadeInSurface.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert fade in surface to texture")
		return
	}
	fadeOutTexture, err := window.ConvertSurfaceToTexture(fadeOutSurface.ID)
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

	// Run event loop
	engine.EventLoop(func(event sdl.Event) bool {
		err = window.Reset(nil)
		if err != nil {
			log.Error().Err(err).Msg("Failed to reset window")
			return false
		}

		switch eventType := event.(type) {
		// Key pressed
		case *sdl.KeyboardEvent:
			switch eventType.Keysym.Sym {
			case sdl.K_UP:
				alpha = uint8(lo.Min([]int16{255, int16(alpha) + 32}))
			case sdl.K_DOWN:
				alpha = uint8(lo.Max([]int16{0, int16(alpha) - 32}))
			}
		}

		// Set fade out texture alpha
		err = fadeOutTexture.SetAlphaMod(alpha)
		if err != nil {
			log.Error().Err(err).Msg("Failed to set fade out texture alpha mod")
			return false
		}

		// Render
		err = window.RenderTexture(&fadeInTexture, 0, 0)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render fade in texture")
			return false
		}
		err = window.RenderTexture(&fadeOutTexture, 0, 0)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render fade out texture")
			return false
		}
		window.Present()

		return true
	})
}
