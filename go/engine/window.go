package engine

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	*sdl.Window
	surface       *sdl.Surface
	renderer      *sdl.Renderer
	surfaceMap    map[uint64]*Surface
	lastSurfaceID uint64
	textureMap    map[uint64]*Texture
	lastTextureID uint64
}

// Create window
func CreateWindow() (win Window, err error) {
	win.Window, err = sdl.CreateWindow("SDL Tutorial", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 480, sdl.WINDOW_SHOWN)
	if err != nil {
		return win, errors.Wrap(err, "failed to create window")
	}

	win.surface, err = win.Window.GetSurface()
	if err != nil {
		return win, errors.Wrap(err, "failed to get window inner surface")
	}

	win.renderer, err = win.Window.GetRenderer()
	if err != nil {
		return win, errors.Wrap(err, "failed to get window inner renderer")
	}

	win.surfaceMap = map[uint64]*Surface{}
	win.textureMap = map[uint64]*Texture{}

	return win, nil
}

// Destroy window
func (w *Window) Destroy() {
	// Destroy surfaces
	for id := range w.surfaceMap {
		err := w.DestroySurface(id)
		if err != nil {
			log.Warn().Err(err).Uint64("id", id).Msg("Failed to destory surface")
		}
	}

	// Destroy textures
	for id := range w.textureMap {
		err := w.DestroyTexture(id)
		if err != nil {
			log.Warn().Err(err).Uint64("id", id).Msg("Failed to destory texture")
		}
	}
}
