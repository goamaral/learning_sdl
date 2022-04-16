package engine

import "github.com/veandco/go-sdl2/sdl"

type SpriteMap struct {
	ID      uint32
	texture *Texture
	spriteW int32
	spriteH int32
}

type Sprite struct {
	texture *Texture
	W       int32
	H       int32
	srcCrop sdl.Rect
}

// Get sprite from sprite map
func (sm *SpriteMap) GetSprite(x uint32, y uint32) Sprite {
	return Sprite{
		texture: sm.texture,
		W:       sm.spriteW,
		H:       sm.spriteH,
		srcCrop: sdl.Rect{
			X: int32(x) * sm.spriteW,
			Y: int32(y) * sm.spriteH,
			W: sm.spriteW,
			H: sm.spriteH,
		},
	}
}
