package ui

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

// Selector is used to show that something has been selected by drawing a rectangle behind it.
type Selector struct {
	position  pixel.Vec
	size      pixel.Rect
	image     *pixel.Sprite
	selected  bool
	fill      *imdraw.IMDraw
	fillColor color.Color
	action    func(selector *Selector)
}

// NewSelector creates a new selector.
func NewSelector(position pixel.Vec, image *pixel.Sprite, action func(selector *Selector)) *Selector {
	selector := &Selector{
		position:  position,
		size:      pixel.R(0, 0, 50, 50),
		image:     image,
		fillColor: pixel.RGB(0, 0.45, 1),
		action:    action,
	}

	selector.SetPosition(position)

	return selector
}

// SetPosition sets the position of the selector.
func (s *Selector) SetPosition(position pixel.Vec) {
	s.position = position
	s.redraw()
}

// Selected sets the state as selected or not selected.
func (s *Selector) Selected(selected bool) {
	s.selected = selected
}

// Position returns the position of the selector.
func (s *Selector) Position() pixel.Vec {
	return s.position
}

// Size returns the size of the selector.
func (s *Selector) Size() pixel.Rect {
	return s.size
}

// Draw the selector to the screen.
func (s *Selector) Draw(t pixel.Target) {
	if s.selected {
		s.fill.Draw(t)
	}

	pos := pixel.V(s.position.X+s.size.W()/2, s.position.Y+s.size.H()/2)
	s.image.Draw(t, pixel.IM.Moved(pos).Scaled(pos, 4))
}

// Update the selector and perform an action if it has been clicked.
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

// redraw will set the properties necessary to draw the selector at a different position.
func (s *Selector) redraw() {
	s.fill = NewRectangle(s.position, s.position.Add(pixel.V(s.size.W(), s.size.H())), s.fillColor)
}
