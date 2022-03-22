package engine

import "github.com/veandco/go-sdl2/sdl"

type Texture struct {
	*sdl.Texture
	ID uint64
}
