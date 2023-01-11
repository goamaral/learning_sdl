/*
	Clip Rendering and Sprite Sheets
	https://lazyfoo.net/tutorials/SDL/11_clip_rendering_and_sprite_sheets/index.php
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
	surface, err := window.LoadSurface("../../../resources/images/dots.png")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load surface")
		return
	}

	// Set player surface transparent color
	surface.SetTransparentColor(engine.ColorByName[engine.ColorName_CYAN])

	// Convert surface to texture
	texture, err := window.Renderer.SurfaceToTexture(surface)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert surface to texture")
		return
	}

	// Create sprite map from texture and get sprites
	spriteMap := engine.NewSpriteMap(texture, 100, 100)
	topLeftSprite := spriteMap.GetSprite(0)
	topRightSprite := spriteMap.GetSprite(1)
	bottomLeftSprite := spriteMap.GetSprite(2)
	bottomRightSprite := spriteMap.GetSprite(3)

	engine.EventLoop(func(getEvent func() sdl.Event) bool {
		err = window.Renderer.Reset()
		if err != nil {
			log.Error().Err(err).Msg("Failed to reset renderer")
			return false
		}

		// Render sprites
		err = window.Renderer.RenderSprite(engine.RenderContext{}, &topLeftSprite, 0, 0, false)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render top left sprite")
			return false
		}
		err = window.Renderer.RenderSprite(engine.RenderContext{}, &topRightSprite, window.W-topRightSprite.W, 0, false)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render top right sprite")
			return false
		}
		err = window.Renderer.RenderSprite(engine.RenderContext{}, &bottomLeftSprite, 0, window.H-bottomLeftSprite.H, false)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render bottom left sprite")
			return false
		}
		err = window.Renderer.RenderSprite(engine.RenderContext{}, &bottomRightSprite, window.W-topRightSprite.W, window.H-bottomLeftSprite.H, false)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render bottom right sprite")
			return false
		}

		window.Renderer.Present()
		time.Sleep(2 * time.Second)

		return false
	})
}
