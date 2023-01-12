/*
	Geometry Rendering
	https://lazyfoo.net/tutorials/SDL/08_geometry_rendering/index.php
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

	engine.EventLoop(func(getEvent func() sdl.Event) bool {
		err = window.Renderer.Reset()
		if err != nil {
			log.Error().Err(err).Msg("Failed to reset window")
			return false
		}

		// Render red rectangle fill
		err = window.Renderer.RenderRectangle(
			engine.RenderContext{DrawColor: &engine.Color_RED},
			(engine.Rectangle)(sdl.Rect{X: window.W / 4, Y: window.H / 4, W: window.W / 2, H: window.H / 2}),
			engine.RenderGeometryMode_FILL,
		)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render red rectangle fill")
			return false
		}

		// Render green rectangle outline
		err = window.Renderer.RenderRectangle(
			engine.RenderContext{DrawColor: &engine.Color_GREEN},
			(engine.Rectangle)(sdl.Rect{X: window.W / 6, Y: window.H / 6, W: window.W * 2 / 3, H: window.H * 2 / 3}),
			engine.RenderGeometryMode_OUTLINE,
		)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render green rectangle outline")
			return false
		}

		// Render horizontal blue line
		err = window.Renderer.RenderLine(
			engine.RenderContext{DrawColor: &engine.Color_BLUE},
			0, window.H/2, window.W, window.H/2,
		)
		if err != nil {
			log.Error().Err(err).Msg("Failed to render horizontal blue line")
			return false
		}

		// Render yellow dots to create intermittent vertical line
		var i int32 = 0
		for ; i < window.H; i += 4 {
			err = window.Renderer.RenderPoint(engine.RenderContext{DrawColor: &engine.Color_YELLOW}, window.W/2, i)
			if err != nil {
				log.Error().Err(err).Msg("Failed to render yellow dots to create intermittent vertical line")
				return false
			}
		}

		window.Renderer.Present()
		time.Sleep(2 * time.Second)

		return false
	})
}
