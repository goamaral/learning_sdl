package engine

type RenderOption interface {
	Apply(renderer *Renderer) error
	Name() RenderOptionName
}

type RenderOptionName string

const (
	RenderOptionName_SetBackgroundColor RenderOptionName = "RenderOption_SetBackgroundColor"
)

func (r *Renderer) ApplyOptions(opts ...RenderOption) (map[RenderOptionName]struct{}, error) {
	appliedOpts := map[RenderOptionName]struct{}{}

	for _, opt := range opts {
		err := opt.Apply(r)
		if err != nil {
			return appliedOpts, err
		}
		appliedOpts[opt.Name()] = struct{}{}
	}

	return appliedOpts, nil
}

type RenderOption_SetBackgroundColor Color

func (opt RenderOption_SetBackgroundColor) Apply(renderer *Renderer) error {
	return renderer.setDrawColor(Color(opt))
}

func (opt RenderOption_SetBackgroundColor) Name() RenderOptionName {
	return RenderOptionName_SetBackgroundColor
}
