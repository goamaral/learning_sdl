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

	color := engine.COLOR_WHITE

	// Run event loop
	engine.EventLoop(func(event sdl.Event) bool {
		switch eventType := event.(type) {
		// Key pressed
		case *sdl.KeyboardEvent:
			switch eventType.Keysym.Sym {
			case sdl.K_q:
				color.Red = uint8(lo.Min([]int16{255, int16(color.Red) + 32}))

			case sdl.K_a:
				color.Red = uint8(lo.Max([]int16{0, int16(color.Red) - 32}))

			case sdl.K_w:
				color.Green = uint8(lo.Min([]int16{255, int16(color.Green) + 32}))

			case sdl.K_s:
				color.Green = uint8(lo.Max([]int16{0, int16(color.Green) - 32}))

			case sdl.K_e:
				color.Blue = uint8(lo.Min([]int16{255, int16(color.Blue) + 32}))

			case sdl.K_d:
				color.Blue = uint8(lo.Max([]int16{0, int16(color.Blue) - 32}))
			}

			err := texture.SetColorModulation(color)
			if err != nil {
				log.Error().Err(err).Msg("Failed to set color modulation")
				return false
			}
		}

		err = window.RenderTexture(&texture, 0, 0)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render texture")
			return false
		}
		window.Present()

		return true
	})
}
