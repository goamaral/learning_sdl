package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

var viewports map[uint32]*Viewport = map[uint32]*Viewport{}
var nextViewportID uint32

type Viewport struct {
	sdl.Rect
	id        uint32
	destroyed bool
}

/* PRIVATE */
func (v *Viewport) save() *Viewport {
	v.id = nextViewportID
	nextViewportID += 1
	viewports[v.id] = v
	return v
}

func destroyViewports() {
	for _, viewport := range viewports {
		viewport.Destroy()
	}
}

/* PUBLIC */
func NewViewportFromSdlRect(sdlRect sdl.Rect) *Viewport {
	return (&Viewport{Rect: sdlRect}).save()
}

func NewViewport(x int32, y int32, w int32, h int32) *Viewport {
	return NewViewportFromSdlRect(sdl.Rect{X: x, Y: y, W: w, H: h})
}

func (v *Viewport) Destroy() {
	if !v.destroyed {
		v.destroyed = true
		delete(viewports, v.id)
	}
}
