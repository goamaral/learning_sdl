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
	backgroundSurfaceId, err := window.LoadSurface("../../../resources/images/background.png")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load background surface")
		return
	}
	playerSurfaceId, err := window.LoadSurface("../../../resources/images/player.png")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load player surface")
		return
	}
	playerSurface, err := window.GetSurface(playerSurfaceId)
	if err != nil {
		log.Error().Err(err).Msg("failed to get player surface")
	}

	// Set player surface color key
	playerSurface.SetColorKey(engine.COLOR_CYAN)

	// Convert surfaces to textures
	backgroundTextureId, err := window.ConvertSurfaceToTexture(backgroundSurfaceId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert background surface to texture")
		return
	}
	playerTextureId, err := window.ConvertSurfaceToTexture(playerSurfaceId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert player surface to texture")
		return
	}

	// Reset window
	err = window.Reset(&engine.COLOR_BLACK)
	if err != nil {
		log.Error().Err(err).Msg("Failed to reset window")
		return
	}

	// Render texture to every viewport
	window.RenderTexture(backgroundTextureId, 0, nil)
	window.RenderTexture(playerTextureId, 0, &sdl.Rect{X: 240, Y: 190})

	// Present
	window.Present()
	engine.ProcessEvents(nil)
	sdl.Delay(2000)
}
