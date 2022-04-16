package engine

import "github.com/veandco/go-sdl2/sdl"

type Texture struct {
	*sdl.Texture
	ID uint32
	W  int32
	H  int32
}
