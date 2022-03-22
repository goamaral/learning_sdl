package engine

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

// Save viewport
func (w *Window) SaveViewport(viewport *Viewport) {
	viewport.ID = w.GetNextResourceID()
	w.viewportMap[viewport.ID] = *viewport
}

// Get viewport
func (w *Window) GetViewport(id uint64) (Viewport, error) {
	viewport, exists := w.viewportMap[id]
	if !exists {
		return viewport, errors.Errorf("viewport %d not found", id)
	}

	return viewport, nil
}

// Create viewport
func (w *Window) CreateViewport(x, y, width, height int32) uint64 {
	viewport := Viewport{Rect: sdl.Rect{X: x, Y: y, W: width, H: height}}
	w.SaveViewport(&viewport)

	return viewport.ID
}
