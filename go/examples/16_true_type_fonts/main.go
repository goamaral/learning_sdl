/*
	True type fonts
	https://lazyfoo.net/tutorials/SDL/16_true_type_fonts/index.php
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
	window, err := engine.NewWindow(640, 480, true)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create window")
		return
	}

	// Load font
	font, err := engine.LoadFont("../../../resources/fonts/lazy.ttf", 28)
	if err != nil {
		log.Error().Err(err).Msg("Failed to load font")
		return
	}

	// Text to texture
	texture, err := window.Renderer.TextToTexture("The quick brown fox jumps over the lazy dog", font, engine.Color_BLACK)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create texture from text")
		return
	}

	engine.EventLoop(func(getEvent func() sdl.Event) bool {
		// Reset renderer
		err = window.Renderer.Reset(engine.RenderOption_SetBackgroundColor(engine.Color_WHITE))
		if err != nil {
			log.Error().Err(err).Msg("Failed to reset renderer")
			return false
		}

		// Render texture
		err = window.Renderer.RenderTexture(
			engine.RenderContext{},
			texture,
			(window.W-texture.W)/2, (window.H-texture.H)/2,
			nil,
		)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render texture")
			return false
		}

		window.Renderer.Present()
		time.Sleep(2 * time.Second)

		return false
	})
}
