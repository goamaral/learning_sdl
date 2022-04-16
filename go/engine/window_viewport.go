package engine

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

// Save viewport
func (w *Window) SaveViewport(viewport *Viewport) {
	viewport.ID = w.GetNextResourceID()
	w.viewports[viewport.ID] = *viewport
}

// Get viewport
func (w *Window) GetViewport(id uint32) (Viewport, error) {
	viewport, exists := w.viewports[id]
	if !exists {
		return viewport, errors.Errorf("viewport %d not found", id)
	}

	return viewport, nil
}

// Create viewport
func (w *Window) CreateViewport(x, y, width, height int32) Viewport {
	viewport := Viewport{
		Rect:   sdl.Rect{X: x, Y: y, W: width, H: height},
		window: w,
	}
	w.SaveViewport(&viewport)

	return viewport
}
