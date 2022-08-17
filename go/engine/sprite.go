package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type SpriteMap struct {
	ID         uint32
	sdlTexture *Texture
	SpriteW    int32
	SpriteH    int32
}

// Get sprite from sprite map
func (sm *SpriteMap) GetSprite(index uint) Sprite {
	return Sprite{
		Texture: sm.sdlTexture,
		W:       sm.SpriteW,
		H:       sm.SpriteH,
		srcCrop: sdl.Rect{
			X: (int32(index) * sm.SpriteW) % sm.sdlTexture.W,
			Y: ((int32(index) * sm.SpriteW) / sm.sdlTexture.W) * sm.SpriteH,
			W: int32(sm.SpriteW),
			H: int32(sm.SpriteH),
		},
	}
}

// Get sprite map length
func (sm *SpriteMap) Len() uint {
	return uint((sm.sdlTexture.W / sm.SpriteW) * (sm.sdlTexture.H / sm.SpriteH))
}

type Sprite struct {
	W       int32
	H       int32
	Texture *Texture // TODO: Make private
	srcCrop sdl.Rect
}
