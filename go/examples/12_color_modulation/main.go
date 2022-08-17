/*
	Color Modulation
	https://lazyfoo.net/tutorials/SDL/12_color_modulation/index.php
*/

package main

import (
	"github.com/Goamaral/learning_sdl/engine"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
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
	surface, err := window.LoadSurface("../../../resources/images/colors.png")
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

	color := engine.ColorByName[engine.ColorName_WHITE]

	engine.EventLoop(func(getEvent func() sdl.Event) bool {
		err = window.Renderer.Reset()
		if err != nil {
			log.Error().Err(err).Msg("Failed to reset renderer")
			return false
		}

		for sdlEvent := getEvent(); sdlEvent != nil; sdlEvent = getEvent() {
			switch event := sdlEvent.(type) {
			// Key pressed
			case *sdl.KeyboardEvent:
				switch event.Keysym.Sym {
				case sdl.K_q:
					color.R = uint8(lo.Min([]int16{255, int16(color.R) + 32}))

				case sdl.K_a:
					color.R = uint8(lo.Max([]int16{0, int16(color.R) - 32}))

				case sdl.K_w:
					color.G = uint8(lo.Min([]int16{255, int16(color.G) + 32}))

				case sdl.K_s:
					color.G = uint8(lo.Max([]int16{0, int16(color.G) - 32}))

				case sdl.K_e:
					color.B = uint8(lo.Min([]int16{255, int16(color.B) + 32}))

				case sdl.K_d:
					color.B = uint8(lo.Max([]int16{0, int16(color.B) - 32}))
				}
			}
		}

		// Set color modulation
		err := texture.SetColorModulation(color)
		if err != nil {
			log.Error().Err(err).Msg("Failed to set color modulation")
			return false
		}

		// Render texture
		err = window.Renderer.RenderTexture(engine.RenderContext{}, &texture, 0, 0, engine.TextureRenderMode_DEFAULT)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render texture")
			return false
		}

		window.Renderer.Present()

		return true
	})
}
