package engine

import "github.com/veandco/go-sdl2/sdl"

type GeometryRenderMode uint8

const (
	GEOMETRY_RENDER_MODE_FILL GeometryRenderMode = iota
	GEOMETRY_RENDER_MODE_OUTLINE
)

type Rectangle struct {
	sdl.Rect
	Color      Color
	RenderMode GeometryRenderMode
}
