package engine

import "github.com/veandco/go-sdl2/sdl"

// Save sprite map
func (w *Window) SaveSpriteMap(spriteMap *SpriteMap) {
	spriteMap.ID = w.GetNextResourceID()
	w.spriteMaps[spriteMap.ID] = *spriteMap
}

// Create sprite map
func (w *Window) CreateSpriteMap(texture *Texture, spriteW int32, spriteH int32) SpriteMap {
	spriteMap := SpriteMap{texture: texture, spriteW: spriteW, spriteH: spriteH}
	w.SaveSpriteMap(&spriteMap)

	return spriteMap
}

// Render sprite
func (w *Window) RenderSprite(sprite *Sprite, x int32, y int32) error {
	return w.renderer.Copy(
		sprite.texture.Texture,
		&sprite.srcCrop,
		&sdl.Rect{X: x, Y: y, W: sprite.W, H: sprite.H},
	)
}
