package engine

import "github.com/veandco/go-sdl2/sdl"

type Viewport struct {
	sdl.Rect
	ID uint64
}
