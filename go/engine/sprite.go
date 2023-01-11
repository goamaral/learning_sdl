package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type SpriteMap struct {
	Texture *Texture
	SpriteW int32
	SpriteH int32
}

type Sprite struct {
	srcCrop sdl.Rect

	Texture *Texture
	W       int32
	H       int32
}

func NewSpriteMap(texture *Texture, spriteW int32, spriteH int32) SpriteMap {
	return SpriteMap{Texture: texture, SpriteW: spriteW, SpriteH: spriteH}
}

func (sm *SpriteMap) GetSprite(index uint) Sprite {
	return Sprite{
		Texture: sm.Texture,
		W:       sm.SpriteW,
		H:       sm.SpriteH,
		srcCrop: sdl.Rect{
			X: (int32(index) * sm.SpriteW) % sm.Texture.W,
			Y: ((int32(index) * sm.SpriteW) / sm.Texture.W) * sm.SpriteH,
			W: int32(sm.SpriteW),
			H: int32(sm.SpriteH),
		},
	}
}

func (sm *SpriteMap) Len() uint {
	return uint((sm.Texture.W / sm.SpriteW) * (sm.Texture.H / sm.SpriteH))
}
