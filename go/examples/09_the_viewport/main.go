/*
	The Viewport
	https://lazyfoo.net/tutorials/SDL/09_the_viewport/index.php
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

	// Load surface
	surface, err := window.LoadSurface("../../../resources/images/texture.png")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load surface")
		return
	}

	// Convert surface to texture
	texture, err := window.Renderer.SurfaceToTexture(surface)
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

	engine.EventLoop(func(getEvent func() sdl.Event) bool {
		err = window.Renderer.Reset()
		if err != nil {
			log.Error().Err(err).Msg("Failed to reset window")
			return false
		}

		// Render texture to top left viewport
		err = window.Renderer.RenderTexture(
			engine.RenderContext{Viewport: topLeftViewport},
			texture,
			0, 0,
			&engine.RenderTextureOptions{Stretch: true},
		)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render texture to top left viewport")
			return false
		}

		// Render texture to top right viewport
		window.Renderer.RenderTexture(
			engine.RenderContext{Viewport: topRightViewport},
			texture,
			0, 0,
			&engine.RenderTextureOptions{Stretch: true},
		)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render texture to top right viewport")
			return false
		}

		// Render texture to bottom viewport
		window.Renderer.RenderTexture(
			engine.RenderContext{Viewport: bottomViewport},
			texture,
			0, 0,
			&engine.RenderTextureOptions{Stretch: true},
		)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render texture to bottom viewport")
			return false
		}

		window.Renderer.Present()
		time.Sleep(2 * time.Second)

		return false
	})
}
