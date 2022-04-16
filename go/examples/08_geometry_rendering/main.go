/*
	Geometry Rendering
	https://lazyfoo.net/tutorials/SDL/08_geometry_rendering/index.php
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

	// Reset window
	err = window.Reset(&engine.COLOR_BLACK)
	if err != nil {
		log.Error().Err(err).Msg("Failed to reset window")
		return
	}

	// Render red rectangle fill
	err = window.RenderRectangle(&engine.Rectangle{
		Rect:       sdl.Rect{X: window.W / 4, Y: window.H / 4, W: window.W / 2, H: window.H / 2},
		Color:      engine.COLOR_RED,
		RenderMode: engine.GEOMETRY_RENDER_MODE_FILL,
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to render red rectangle fill")
		return
	}

	// Render green rectangle outline
	err = window.RenderRectangle(&engine.Rectangle{
		Rect:       sdl.Rect{X: window.W / 6, Y: window.H / 6, W: window.W * 2 / 3, H: window.H * 2 / 3},
		Color:      engine.COLOR_GREEN,
		RenderMode: engine.GEOMETRY_RENDER_MODE_OUTLINE,
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to render green rectangle outline")
		return
	}

	// Render horizontal blue line
	err = window.RenderLine(0, window.H/2, window.W, window.H/2, &engine.COLOR_BLUE)
	if err != nil {
		log.Error().Err(err).Msg("Failed to render horizontal blue line")
		return
	}

	// Render yellow dots to create intermittent vertical line
	var i int32 = 0
	for ; i < window.H; i += 4 {
		err = window.RenderPoint(window.W/2, i, &engine.COLOR_YELLOW)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render yellow dots to create intermittent vertical line")
			return
		}
	}

	// Present
	window.Present()
	engine.ProcessEvents(nil)
	sdl.Delay(2000)
}
