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
	window, err := engine.CreateWindow()
	if err != nil {
		log.Error().Err(err).Msg("Failed to create window")
		return
	}
	defer window.Destroy()

	// Load surface
	walkingSurface, err := window.LoadSurface("../../../resources/images/walking.png")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load walking in surface")
		return
	}

	// Convert surface to texture
	walkingTexture, err := window.ConvertSurfaceToTexture(walkingSurface)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert walking surface to texture")
		return
	}

	// Create sprite map
	spriteMap := window.CreateSpriteMap(&walkingTexture, 64, 205)

	// Run event loop
	var currentFrame uint
	engine.EventLoop(func(getEvent func() sdl.Event) bool {
		err = window.Renderer.Reset()
		if err != nil {
			log.Error().Err(err).Msg("Failed to reset renderer")
			return false
		}

		// Process event
		for getEvent() != nil {
		}

		// Get current frame sprite
		sprite := spriteMap.GetSprite(currentFrame % spriteMap.Len())

		// Render sprite
		err = window.Renderer.RenderSprite(engine.RenderContext{}, &sprite, 0, 0, engine.TextureRenderMode_DEFAULT)
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
