package engine

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	*sdl.Window
	W        int32
	H        int32
	surface  *sdl.Surface
	renderer *sdl.Renderer

	/* Resources */
	lastResourceID uint32
	surfaces       map[uint32]Surface
	textures       map[uint32]Texture
	viewports      map[uint32]Viewport // Do not need to be deallocated
	spriteMaps     map[uint32]SpriteMap
}

// Create window
func CreateWindow() (Window, error) {
	var err error

	win := Window{
		W:          640,
		H:          480,
		surfaces:   map[uint32]Surface{},
		textures:   map[uint32]Texture{},
		viewports:  map[uint32]Viewport{},
		spriteMaps: map[uint32]SpriteMap{},
	}

	win.Window, err = sdl.CreateWindow("SDL Tutorial", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, win.W, win.H, sdl.WINDOW_SHOWN)
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
	for id := range w.surfaces {
		err := w.DestroySurface(id)
		if err != nil {
			log.Warn().Err(err).Uint32("id", id).Msg("Failed to destory surface")
		}
	}

	// Destroy textures
	for id := range w.textures {
		err := w.DestroyTexture(id)
		if err != nil {
			log.Warn().Err(err).Uint32("id", id).Msg("Failed to destory texture")
		}
	}
}

// Get next resource id
// FUTURE: Make it thread safe
func (w *Window) GetNextResourceID() uint32 {
	nextID := w.lastResourceID + 1
	w.lastResourceID = nextID

	return nextID
}
