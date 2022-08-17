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
	surface  *sdl.Surface // TODO: Make it *Surface
	Renderer *Renderer

	/* Resources */
	lastResourceID uint32
	textures       map[uint32]Texture
	spriteMaps     map[uint32]SpriteMap
}

func CreateWindow(useVsync bool) (Window, error) {
	var err error

	win := Window{
		W:          640,
		H:          480,
		textures:   map[uint32]Texture{},
		spriteMaps: map[uint32]SpriteMap{},
	}

	win.Window, err = sdl.CreateWindow("SDL Tutorial", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, win.W, win.H, sdl.WINDOW_SHOWN)
	if err != nil {
		return win, errors.Wrap(err, "failed to create window")
	}

	sdlRenderer, err := win.Window.GetRenderer()
	if err != nil {
		return win, errors.Wrap(err, "failed to get window renderer")
	}
	if useVsync {
		sdlRenderer, err = sdl.CreateRenderer(win.Window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
		if err != nil {
			return win, errors.Wrap(err, "failed to create renderer")
		}
	}
	win.Renderer = NewRenderer(sdlRenderer)

	win.surface, err = win.Window.GetSurface()
	if err != nil {
		return win, errors.Wrap(err, "failed to get window surface")
	}

	return win, nil
}

func (w *Window) Destroy() {
	// Destroy textures
	for id := range w.textures {
		err := w.DestroyTexture(id)
		if err != nil {
			log.Warn().Err(err).Uint32("id", id).Msg("Failed to destory texture")
		}
	}
}

// FUTURE: Make it thread safe
func (w *Window) GetNextResourceID() uint32 {
	nextID := w.lastResourceID + 1
	w.lastResourceID = nextID

	return nextID
}

func (w *Window) GetPixelFormat() *sdl.PixelFormat {
	return w.surface.Format
}
