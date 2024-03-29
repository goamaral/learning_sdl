/*
	Animated Sprites and VSync
	https://lazyfoo.net/tutorials/SDL/14_animated_sprites_and_vsync/index.php
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

	// Load surface
	walkingSurface, err := window.LoadSurface("../../../resources/images/walking.png")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load walking in surface")
		return
	}

	// Convert surface to texture
	walkingTexture, err := window.Renderer.SurfaceToTexture(walkingSurface)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert walking surface to texture")
		return
	}

	// Create sprite map
	spriteMap := engine.NewSpriteMap(walkingTexture, 64, 205)

	var currentFrame uint
	x := (window.W - spriteMap.SpriteW) / 2
	y := (window.H - spriteMap.SpriteH) / 2

	engine.EventLoop(func(getEvent func() sdl.Event) bool {
		// Process events
		for event := getEvent(); event != nil; event = getEvent() {
			switch eventType := event.(type) {
			// Key pressed
			case *sdl.KeyboardEvent:
				switch eventType.Keysym.Sym {
				case sdl.K_ESCAPE:
					return false
				}
			}
		}

		// Reset renderer
		err = window.Renderer.Reset()
		if err != nil {
			log.Error().Err(err).Msg("Failed to reset renderer")
			return false
		}

		// Get current frame sprite
		sprite := spriteMap.GetSprite(currentFrame % spriteMap.Len())

		// Render sprite
		err = window.Renderer.RenderSprite(engine.RenderContext{}, &sprite, x, y, false)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render sprite")
			return false
		}

		window.Renderer.Present()
		currentFrame += 1
		time.Sleep(time.Second / 10)

		return true
	})
}
