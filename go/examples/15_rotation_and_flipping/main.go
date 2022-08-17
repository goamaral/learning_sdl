/*
	Rotation and Flipping
	https://lazyfoo.net/tutorials/SDL/15_rotation_and_flipping/index.php
*/

package main

import (
	"github.com/Goamaral/learning_sdl/engine"
	"github.com/rs/zerolog/log"
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
	window, err := engine.CreateWindow(true)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create window")
		return
	}
	defer window.Destroy()

	// Load surface
	arrowSurface, err := window.LoadSurface("../../../resources/images/arrow.png")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load arrow in surface")
		return
	}

	// Convert surface to texture
	arrowTexture, err := window.ConvertSurfaceToTexture(arrowSurface)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert arrow surface to texture")
		return
	}

	var angle float64
	flip := engine.TextureRenderFlip_NONE
	x := (window.W - arrowTexture.W) / 2
	y := (window.H - arrowTexture.H) / 2

	engine.EventLoop(func(getEvent func() sdl.Event) bool {
		for sdlEvent := getEvent(); sdlEvent != nil; sdlEvent = getEvent() {
			switch eventType := sdlEvent.(type) {
			// Key pressed
			case *sdl.KeyboardEvent:
				switch eventType.Keysym.Sym {
				case sdl.K_a:
					angle -= 60
				case sdl.K_d:
					angle += 60
				case sdl.K_q:
					flip = engine.TextureRenderFlip_HORIZONTAL
				case sdl.K_w:
					flip = engine.TextureRenderFlip_NONE
				case sdl.K_e:
					flip = engine.TextureRenderFlip_VERTICAL
				}
			}
		}

		// Reset renderer
		err = window.Renderer.Reset()
		if err != nil {
			log.Error().Err(err).Msg("Failed to reset renderer")
			return false
		}

		// Render texture
		err = window.Renderer.RenderTexture(
			engine.RenderContext{},
			&arrowTexture,
			x, y,
			&engine.RenderTextureOptions{Flip: flip, Angle: angle},
		)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render texture")
			return false
		}

		window.Renderer.Present()

		return true
	})
}
