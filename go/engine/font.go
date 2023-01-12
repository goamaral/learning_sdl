package engine

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/ttf"
)

type Font struct {
	abstractResource

	sdlFont *ttf.Font
}

/* PRIVATE */
var fontManager = newResourceManager()

func newFont(sdlFont *ttf.Font) *Font {
	return fontManager.Save(&Font{sdlFont: sdlFont}).(*Font)
}

/* PUBLIC */
func (f *Font) Destroy() {
	if !f.destroyed {
		f.destroyed = true
		delete(fontManager.resources, f.id)
		f.sdlFont.Close()
	}
}

// Supports ttf
func LoadFont(path string, size int) (*Font, error) {
	pathParts := strings.Split(path, ".")
	extension := pathParts[len(pathParts)-1]
	var sdlFont *ttf.Font
	var err error

	switch extension {
	case "ttf":
		sdlFont, err = ttf.OpenFont(path, size)
		if err != nil {
			return nil, errors.Wrap(err, "failed to load ttf font")
		}
	default:
		return nil, errors.Errorf("unsupported font file format %s")
	}

	return newFont(sdlFont), nil
}
