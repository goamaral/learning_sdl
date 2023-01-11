package engine

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type Renderer struct {
	abstractResource
	sdlRenderer *sdl.Renderer

	defaultViewport  *Viewport
	defaultDrawColor Color
}

type RenderContext struct {
	Viewport  *Viewport
	DrawColor *Color
}

type RenderTextureFlip sdl.RendererFlip

const (
	RenderTextureFlip_NONE       = RenderTextureFlip(sdl.FLIP_NONE)
	RenderTextureFlip_HORIZONTAL = RenderTextureFlip(sdl.FLIP_HORIZONTAL)
	RenderTextureFlip_VERTICAL   = RenderTextureFlip(sdl.FLIP_VERTICAL)
)

type RenderTextureOptions struct {
	Stretch bool
	Flip    RenderTextureFlip
	Angle   float64
}

type RenderGeometryMode uint

const (
	RenderGeometryMode_DEFAULT RenderGeometryMode = iota
	RenderGeometryMode_FILL
	RenderGeometryMode_OUTLINE
)

/* PRIVATE */
var rendererManager = newResourceManager()

func newRenderer(sdlRenderer *sdl.Renderer) *Renderer {
	return rendererManager.Save(&Renderer{
		sdlRenderer:      sdlRenderer,
		defaultViewport:  NewViewportFromSdlRect(sdlRenderer.GetViewport()),
		defaultDrawColor: ColorByName[ColorName_BLACK],
	}).(*Renderer)
}

func (r Renderer) setDrawColor(color Color) error {
	return r.sdlRenderer.SetDrawColor(color.R, color.G, color.B, color.A)
}

func (r Renderer) applyRenderContext(ctx RenderContext) (RenderContext, error) {
	// Set viewport
	if ctx.Viewport == nil {
		ctx.Viewport = r.defaultViewport
	}
	err := r.sdlRenderer.SetViewport((*sdl.Rect)(&ctx.Viewport.Rect))
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
func (r *Renderer) Destroy() {
	if !r.destroyed {
		r.destroyed = true
		delete(rendererManager.resources, r.id)
		r.sdlRenderer.Destroy()
	}
}

func (r Renderer) Reset() error {
	// Reset draw color
	err := r.setDrawColor(r.defaultDrawColor)
	if err != nil {
		return errors.Wrap(err, "failed to reset renderer draw color")
	}

	// Reset viewport
	err = r.sdlRenderer.SetViewport((*sdl.Rect)(&r.defaultViewport.Rect))
	if err != nil {
		return errors.Wrap(err, "failed to reset renderer viewport")
	}

	// Clear renderer
	err = r.sdlRenderer.Clear()
	if err != nil {
		return errors.Wrap(err, "failed to reset renderer")
	}

	return nil
}

func (r Renderer) Present() {
	r.sdlRenderer.Present()
}

func (r Renderer) SurfaceToTexture(surface *Surface) (*Texture, error) {
	sdlTexture, err := r.sdlRenderer.CreateTextureFromSurface(surface.sdlSurface)
	if err != nil {
		return nil, err
	}

	return newTexture(sdlTexture, surface.sdlSurface.W, surface.sdlSurface.H), nil
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
	err = r.sdlRenderer.CopyEx(texture.sdlTexture, nil, &sdl.Rect{X: x, Y: y, W: w, H: h}, opts.Angle, nil, sdl.RendererFlip(opts.Flip))
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
	err = r.sdlRenderer.Copy(sprite.Texture.sdlTexture, &sprite.srcCrop, &sdl.Rect{X: x, Y: y, W: w, H: h})
	if err != nil {
		return errors.Wrap(err, "failed to render texture")
	}

	return nil
}

func (r Renderer) RenderRectangle(ctx RenderContext, rect Rectangle, renderMode RenderGeometryMode) error {
	// Apply render context
	ctx, err := r.applyRenderContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to apply render context")
	}

	// Render rectangle
	switch renderMode {
	case RenderGeometryMode_DEFAULT:
	case RenderGeometryMode_FILL:
		err = r.sdlRenderer.FillRect((*sdl.Rect)(&rect))
	case RenderGeometryMode_OUTLINE:
		err = r.sdlRenderer.DrawRect((*sdl.Rect)(&rect))
	default:
		return errors.Errorf("Invalid geometry render mode %d", renderMode)
	}
	if err != nil {
		return errors.Wrap(err, "failed to render rectangle")
	}

	return nil
}

func (r Renderer) RenderLine(ctx RenderContext, x1, y1, x2, y2 int32) error {
	// Apply render context
	ctx, err := r.applyRenderContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to apply render context")
	}

	// Render line
	err = r.sdlRenderer.DrawLine(x1, y1, x2, y2)
	if err != nil {
		return errors.Wrap(err, "failed to render line")
	}

	return nil
}

func (r Renderer) RenderPoint(ctx RenderContext, x, y int32) error {
	// Apply render context
	ctx, err := r.applyRenderContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to apply render context")
	}

	// Render point
	err = r.sdlRenderer.DrawPoint(x, y)
	if err != nil {
		return errors.Wrap(err, "failed to render line")
	}

	return nil
}
