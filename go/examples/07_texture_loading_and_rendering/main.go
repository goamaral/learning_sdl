/*
	Texture Loading and Rendering
	https://lazyfoo.net/tutorials/SDL/07_texture_loading_and_rendering
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

	// Load surface
	surfaceId, err := window.LoadSurface("../../../resources/images/texture.png")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load surface")
		return
	}

	// Convert surface to texture
	textureId, err := window.ConvertSurfaceToTexture(surfaceId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert surface to texture")
		return
	}

	// Render texture
	err = window.Reset(&engine.COLOR_BLACK)
	if err != nil {
		log.Error().Err(err).Msg("Failed to reset window renderer")
		return
	}
	err = window.RenderTexture(textureId, 0, nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to render texture")
		return
	}

	// Present
	window.Present()
	engine.ProcessEvents(nil)
	sdl.Delay(2000)
}
