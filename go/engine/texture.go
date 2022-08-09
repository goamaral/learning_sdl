package engine

import "github.com/veandco/go-sdl2/sdl"

type Texture struct {
	*sdl.Texture
	ID uint32
	W  int32
	H  int32
}

// Set color modulation
func (t *Texture) SetColorModulation(color Color) error {
	return t.Texture.SetColorMod(color.Red, color.Green, color.Blue)
}
