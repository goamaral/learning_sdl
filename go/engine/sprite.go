package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type SpriteMap struct {
	ID      uint32
	texture *Texture
	spriteW int32
	spriteH int32
}

// Get sprite from sprite map
func (sm *SpriteMap) GetSprite(index uint) Sprite {
	return Sprite{
		Texture: sm.texture,
		W:       sm.spriteW,
		H:       sm.spriteH,
		srcCrop: sdl.Rect{
			X: (int32(index) * sm.spriteW) % sm.texture.W,
			Y: ((int32(index) * sm.spriteW) / sm.texture.W) * sm.spriteH,
			W: int32(sm.spriteW),
			H: int32(sm.spriteH),
		},
	}
}

// Get sprite map length
func (sm *SpriteMap) Len() uint {
	return uint((sm.texture.W / sm.spriteW) * (sm.texture.H / sm.spriteH))
}

type Sprite struct {
	W       int32
	H       int32
	Texture *Texture
	srcCrop sdl.Rect
}
