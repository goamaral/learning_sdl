/*
	Texture Loading and Rendering
	https://lazyfoo.net/tutorials/SDL/07_texture_loading_and_rendering
*/

package main

import (
	"github.com/rs/zerolog/log"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/Goamaral/learning_sdl/engine"
)

func main() {
	err := engine.Init()
	if err != nil {
		log.Error().Err(err).Msg("Failed to initialized engine")
		return
	}
	defer engine.Quit()

	window, err := engine.CreateWindow()
	if err != nil {
		log.Error().Err(err).Msg("Failed to create window")
		return
	}
	defer window.Destroy()

	surface, err := window.LoadSurface("../../../resources/images/texture.png")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load surface")
		return
	}

	texture, err := window.ConvertSurfaceToTexture(surface.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert surface to texture")
		return
	}

	window.Reset(engine.COLOR_BLACK)
	window.RenderTexture(texture.ID)
	window.Present()

	engine.ProcessEvents(nil)
	sdl.Delay(2000)
}
