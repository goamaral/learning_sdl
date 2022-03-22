package engine

var (
	COLOR_BLACK = Color{0, 0, 0, 255}
)

type Color struct {
	red   uint8
	green uint8
	blue  uint8
	alpha uint8
}
