/*
	Color Keying
	https://lazyfoo.net/tutorials/SDL/10_color_keying/index.php
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
	window, err := engine.CreateWindow()
	if err != nil {
		log.Error().Err(err).Msg("Failed to create window")
		return
	}
	defer window.Destroy()

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
	playerSurface.SetTransparentColor(engine.ColorByName[engine.ColorName_CYAN])

	// Convert surfaces to textures
	backgroundTexture, err := window.ConvertSurfaceToTexture(backgroundSurface.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert background surface to texture")
		return
	}
	playerTexture, err := window.ConvertSurfaceToTexture(playerSurface.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert player surface to texture")
		return
	}

	// Reset window
	err = window.Reset(nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to reset window")
		return
	}

	// Render texture to every viewport
	window.RenderTexture(&backgroundTexture, 0, 0)
	window.RenderTexture(&playerTexture, 240, 190)

	// Present
	window.Present()
	engine.ProcessEvents(nil)
	sdl.Delay(2000)
}
