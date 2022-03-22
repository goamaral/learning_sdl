/*
	The Viewport
	https://lazyfoo.net/tutorials/SDL/09_the_viewport/index.php
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

	// Create top left viewport
	topLeftViewportId := window.CreateViewport(0, 0, window.Width/2, window.Height/2)

	// Create top right viewport
	topRightViewportId := window.CreateViewport(window.Width/2, 0, window.Width/2, window.Height/2)

	// Create bottom viewport
	bottomViewportId := window.CreateViewport(0, window.Height/2, window.Width, window.Height/2)

	// Reset window
	err = window.Reset(&engine.COLOR_BLACK)
	if err != nil {
		log.Error().Err(err).Msg("Failed to reset window")
		return
	}

	// Render texture to every viewport
	window.RenderTexture(textureId, topLeftViewportId)
	window.RenderTexture(textureId, topRightViewportId)
	window.RenderTexture(textureId, bottomViewportId)

	// Present
	window.Present()
	engine.ProcessEvents(nil)
	sdl.Delay(2000)
}
