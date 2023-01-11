package engine

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	abstractResource
	sdlWindow *sdl.Window

	W        int32
	H        int32
	Surface  *Surface
	Renderer *Renderer
}

/* PRIVATE */
var windowManager = newResourceManager()

func newWindow(sdlWindow *sdl.Window, W int32, H int32, surface *Surface, renderer *Renderer) *Window {
	return windowManager.Save(&Window{
		sdlWindow: sdlWindow,

		W:        W,
		H:        H,
		Surface:  surface,
		Renderer: renderer,
	}).(*Window)
}

/* PUBLIC */
func NewWindow(W, H int32, useVsync bool) (*Window, error) {
	sdlWindow, err := sdl.CreateWindow("SDL Tutorial", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, W, H, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create window")
	}

	var createRendererFlags uint32 = sdl.RENDERER_ACCELERATED
	if useVsync {
		createRendererFlags |= sdl.RENDERER_PRESENTVSYNC
	}
	sdlRenderer, err := sdl.CreateRenderer(sdlWindow, -1, createRendererFlags)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create window renderer")
	}

	sdlSurface, err := sdlWindow.GetSurface()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get window surface")
	}

	window := newWindow(sdlWindow, W, H, nil, newRenderer(sdlRenderer))
	window.Surface = newSurface(sdlSurface, window)
	window.Surface.destroyed = true

	return window, nil
}

func (w *Window) Destroy() {
	if !w.destroyed {
		w.destroyed = true
		delete(windowManager.resources, w.id)
		w.sdlWindow.Destroy()
	}
}

func (w *Window) GetPixelFormat() *sdl.PixelFormat {
	return w.Surface.GetPixelFormat()
}
