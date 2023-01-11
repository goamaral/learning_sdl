package engine

import "github.com/veandco/go-sdl2/sdl"

type Texture struct {
	abstractResource

	sdlTexture *sdl.Texture
	W          int32
	H          int32
}

/* PRIVATE */
var textureManager = newResourceManager()

func newTexture(sdlTexture *sdl.Texture, W int32, H int32) *Texture {
	return textureManager.Save(&Texture{sdlTexture: sdlTexture, W: W, H: H}).(*Texture)
}

/* PUBLIC */
func (t *Texture) Destroy() {
	if !t.destroyed {
		t.destroyed = true
		delete(textureManager.resources, t.id)
		t.sdlTexture.Destroy()
	}
}

/* SDL COMPATIBILITY */
func (t *Texture) SetColorModulation(color Color) error {
	return t.sdlTexture.SetColorMod(color.R, color.G, color.B)
}

func (t *Texture) SetBlendMode(bm sdl.BlendMode) error {
	return t.sdlTexture.SetBlendMode(bm)
}

func (t *Texture) SetAlphaMod(alpha uint8) error {
	return t.sdlTexture.SetAlphaMod(alpha)
}
