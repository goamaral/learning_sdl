package engine

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	*sdl.Window
	Width    int32
	Height   int32
	surface  *sdl.Surface
	renderer *sdl.Renderer

	/* Storage */
	surfaceMap    map[uint64]Surface
	lastSurfaceID uint64
	textureMap    map[uint64]Texture
	lastTextureID uint64
}

// Create window
func CreateWindow() (Window, error) {
	var err error

	win := Window{
		Width:      640,
		Height:     480,
		surfaceMap: map[uint64]Surface{},
		textureMap: map[uint64]Texture{},
	}

	win.Window, err = sdl.CreateWindow("SDL Tutorial", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, win.Width, win.Height, sdl.WINDOW_SHOWN)
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
