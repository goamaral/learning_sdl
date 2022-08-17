package engine

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type Renderer struct {
	renderer         *sdl.Renderer
	defaultViewport  *Viewport
	defaultDrawColor Color
}

func NewRenderer(nativeRenderer *sdl.Renderer) *Renderer {
	sdlViewport := nativeRenderer.GetViewport()

	return &Renderer{
		renderer:         nativeRenderer,
		defaultViewport:  NewViewportFromSdlRect(sdlViewport),
		defaultDrawColor: ColorByName[ColorName_BLACK],
	}
}

type RenderContext struct {
	Viewport  *Viewport
	DrawColor *Color
}

/* PRIVATE */
// Set draw color
func (r Renderer) setDrawColor(color Color) error {
	return r.renderer.SetDrawColor(color.R, color.G, color.B, color.A)
}

// Apply render context
func (r Renderer) applyRenderContext(ctx RenderContext) (RenderContext, error) {
	// Set viewport
	if ctx.Viewport == nil {
		ctx.Viewport = r.defaultViewport
	}
	err := r.renderer.SetViewport((*sdl.Rect)(&ctx.Viewport.Rect))
	if err != nil {
		return ctx, errors.Wrap(err, "failed to set viewport")
	}

	// Set draw color
	if ctx.DrawColor == nil {
		ctx.DrawColor = &r.defaultDrawColor
	}
	err = r.setDrawColor(*ctx.DrawColor)
	if err != nil {
		return ctx, errors.Wrap(err, "failed to set draw color")
	}

	return ctx, nil
}

/* PUBLIC */
// Reset renderer
func (r Renderer) Reset() error {
	// Reset draw color
	err := r.setDrawColor(r.defaultDrawColor)
	if err != nil {
		return errors.Wrap(err, "failed to reset renderer draw color")
	}

	// Reset viewport
	err = r.renderer.SetViewport((*sdl.Rect)(&r.defaultViewport.Rect))
	if err != nil {
		return errors.Wrap(err, "failed to reset renderer viewport")
	}

	// Clear renderer
	err = r.renderer.Clear()
	if err != nil {
		return errors.Wrap(err, "failed to reset renderer")
	}

	return nil
}

// Present renderer
func (r Renderer) Present() {
	r.renderer.Present()
}

// Create texture from surface
func (r Renderer) CreateTextureFromSurface(surface *Surface) (Texture, error) {
	sdlTexture, err := r.renderer.CreateTextureFromSurface(surface.getSdlSurface())
	if err != nil {
		return Texture{}, err
	}

	return Texture{W: surface.getSdlSurface().W, H: surface.getSdlSurface().H, Texture: sdlTexture}, nil
}

// Render texture
type TextureRenderFlip sdl.RendererFlip

const (
	TextureRenderFlip_NONE       = TextureRenderFlip(sdl.FLIP_NONE)
	TextureRenderFlip_HORIZONTAL = TextureRenderFlip(sdl.FLIP_HORIZONTAL)
	TextureRenderFlip_VERTICAL   = TextureRenderFlip(sdl.FLIP_VERTICAL)
)

type RenderTextureOptions struct {
	Stretch bool
	Flip    TextureRenderFlip
	Angle   float64
}

func (r Renderer) RenderTexture(ctx RenderContext, texture *Texture, x int32, y int32, opts *RenderTextureOptions) error {
	// Apply render context
	ctx, err := r.applyRenderContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to apply render context")
	}

	// Set default options
	if opts == nil {
		opts = &RenderTextureOptions{}
	}

	// Stretch texture
	var w, h int32
	if opts.Stretch {
		w = ctx.Viewport.W
		h = ctx.Viewport.H
	} else {
		w = texture.W
		h = texture.H
	}

	// Render texture
	err = r.renderer.CopyEx(texture.Texture, nil, &sdl.Rect{X: x, Y: y, W: w, H: h}, opts.Angle, nil, sdl.RendererFlip(opts.Flip))
	if err != nil {
		return errors.Wrap(err, "failed to render texture")
	}

	return nil
}

func (r Renderer) RenderSprite(ctx RenderContext, sprite *Sprite, x int32, y int32, strecth bool) error {
	// Apply render context
	ctx, err := r.applyRenderContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to apply render context")
	}

	// Handle render mode
	var w, h int32
	if strecth {
		w = ctx.Viewport.W
		h = ctx.Viewport.H
	} else {
		w = sprite.W
		h = sprite.H
	}

	// Render texture
	err = r.renderer.Copy(sprite.Texture.Texture, &sprite.srcCrop, &sdl.Rect{X: x, Y: y, W: w, H: h})
	if err != nil {
		return errors.Wrap(err, "failed to render texture")
	}

	return nil
}

// Render rectangle
type GeometryRenderMode uint

const (
	GeometryRenderMode_DEFAULT GeometryRenderMode = iota
	GeometryRenderMode_FILL
	GeometryRenderMode_OUTLINE
)

func (r Renderer) RenderRectangle(ctx RenderContext, rect Rectangle, renderMode GeometryRenderMode) error {
	// Apply render context
	ctx, err := r.applyRenderContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to apply render context")
	}

	// Render rectangle
	switch renderMode {
	case GeometryRenderMode_DEFAULT:
	case GeometryRenderMode_FILL:
		err = r.renderer.FillRect((*sdl.Rect)(&rect))
	case GeometryRenderMode_OUTLINE:
		err = r.renderer.DrawRect((*sdl.Rect)(&rect))
	default:
		return errors.Errorf("Invalid geometry render mode %d", renderMode)
	}
	if err != nil {
		return errors.Wrap(err, "failed to render rectangle")
	}

	return nil
}

// Render line
func (r Renderer) RenderLine(ctx RenderContext, x1, y1, x2, y2 int32) error {
	// Apply render context
	ctx, err := r.applyRenderContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to apply render context")
	}

	// Render line
	err = r.renderer.DrawLine(x1, y1, x2, y2)
	if err != nil {
		return errors.Wrap(err, "failed to render line")
	}

	return nil
}

// Render point
func (r Renderer) RenderPoint(ctx RenderContext, x, y int32) error {
	// Apply render context
	ctx, err := r.applyRenderContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to apply render context")
	}

	// Render point
	err = r.renderer.DrawPoint(x, y)
	if err != nil {
		return errors.Wrap(err, "failed to render line")
	}

	return nil
}
