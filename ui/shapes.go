package ui

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

func NewRectangle(min pixel.Vec, max pixel.Vec, color color.Color) *imdraw.IMDraw {
	rect := imdraw.New(nil)

	rect.Color = color
	rect.Push(min)
	rect.Push(pixel.V(max.X, min.Y))
	rect.Push(max)
	rect.Push(pixel.V(min.X, max.Y))
	rect.Polygon(0)

	return rect
}
