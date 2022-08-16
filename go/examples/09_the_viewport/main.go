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
	surface, err := window.LoadSurface("../../../resources/images/texture.png")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load surface")
		return
	}

	// Convert surface to texture
	texture, err := window.ConvertSurfaceToTexture(surface.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert surface to texture")
		return
	}

	// Create top left viewport
	topLeftViewport := engine.NewViewport(0, 0, window.W/2, window.H/2)

	// Create top right viewport
	topRightViewport := engine.NewViewport(window.W/2, 0, window.W/2, window.H/2)

	// Create bottom viewport
	bottomViewport := engine.NewViewport(0, window.H/2, window.W, window.H/2)

	// Reset window
	err = window.Renderer.Reset()
	if err != nil {
		log.Error().Err(err).Msg("Failed to reset window")
		return
	}

	// Render texture to every viewport
	window.Renderer.RenderTexture(
		engine.RenderContext{Viewport: topLeftViewport},
		0, 0,
		&texture,
		engine.TextureRenderMode_STRETCH,
	)
	window.Renderer.RenderTexture(
		engine.RenderContext{Viewport: topRightViewport},
		0, 0,
		&texture,
		engine.TextureRenderMode_STRETCH,
	)
	window.Renderer.RenderTexture(
		engine.RenderContext{Viewport: bottomViewport},
		0, 0,
		&texture,
		engine.TextureRenderMode_STRETCH,
	)

	// Present
	window.Renderer.Present()
	engine.ProcessEvents(nil)
	sdl.Delay(2000)
}
