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
	surface, err := window.LoadSurface("../../../resources/images/texture.png")
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

	// Render texture
	err = window.Renderer.Reset()
	if err != nil {
		log.Error().Err(err).Msg("Failed to reset window")
		return
	}
	err = window.Renderer.RenderTexture(engine.RenderContext{}, 0, 0, &texture, engine.TextureRenderMode_DEFAULT)
	if err != nil {
		log.Error().Err(err).Msg("Failed to render texture")
		return
	}

	// Present
	window.Renderer.Present()
	engine.ProcessEvents(nil)
	sdl.Delay(2000)
}
