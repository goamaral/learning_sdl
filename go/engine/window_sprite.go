package engine

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
