/*
	Color Keying
	https://lazyfoo.net/tutorials/SDL/10_color_keying/index.php
*/

package main

import (
	"time"

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
	window, err := engine.NewWindow(640, 480, false)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create window")
		return
	}

	// Load surfaces
	backgroundSurface, err := window.LoadSurface("../../../resources/images/background.png")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load background surface")
		return
	}
	playerSurface, err := window.LoadSurface("../../../resources/images/player.png")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load player surface")
		return
	}

	// Set player surface transparent color
	playerSurface.SetTransparentColor(engine.Color_CYAN)

	// Convert surfaces to textures
	backgroundTexture, err := window.Renderer.SurfaceToTexture(backgroundSurface)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert background surface to texture")
		return
	}
	playerTexture, err := window.Renderer.SurfaceToTexture(playerSurface)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert player surface to texture")
		return
	}

	engine.EventLoop(func(getEvent func() sdl.Event) bool {
		err = window.Renderer.Reset()
		if err != nil {
			log.Error().Err(err).Msg("Failed to reset renderer")
			return false
		}

		// Render background texture
		err := window.Renderer.RenderTexture(engine.RenderContext{}, backgroundTexture, 0, 0, nil)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render background texture")
			return false
		}

		// Render player texture
		err = window.Renderer.RenderTexture(engine.RenderContext{}, playerTexture, 240, 190, nil)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render player texture")
			return false
		}

		window.Renderer.Present()
		time.Sleep(2 * time.Second)

		return false
	})
}
