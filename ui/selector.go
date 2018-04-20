package ui

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Selector struct {
	position  pixel.Vec
	size      pixel.Rect
	image     *pixel.Sprite
	selected  bool
	fill      *imdraw.IMDraw
	fillColor color.Color
	action    func(selector *Selector)
}

func NewSelector(position pixel.Vec, image *pixel.Sprite, action func(selector *Selector)) *Selector {
	selector := &Selector{
		position:  position,
		size:      pixel.R(0, 0, 50, 50),
		image:     image,
		fillColor: pixel.RGB(0, 0.45, 1),
		action:    action,
	}

	selector.fill = NewRectangle(selector.position, selector.position.Add(pixel.V(selector.size.W(), selector.size.H())), selector.fillColor)

	return selector
}

func (s *Selector) Selected(selected bool) {
	s.selected = selected
}

func (s *Selector) Draw(t pixel.Target) {
	if s.selected {
		s.fill.Draw(t)
	}

	pos := pixel.V(s.position.X+s.size.W()/2, s.position.Y+s.size.H()/2)
	s.image.Draw(t, pixel.IM.Moved(pos).Scaled(pos, 4))
}

func (s *Selector) Update(win *pixelgl.Window) {
	if win.JustPressed(pixelgl.MouseButtonLeft) {
		mpos := win.MousePosition()

		if mpos.X < s.position.X || mpos.X > s.position.X+s.size.W() || mpos.Y < s.position.Y || mpos.Y > s.position.Y+s.size.H() {
			// do nothing
		} else if s.action != nil {
			s.selected = true
			s.action(s)
		}
	}
}
