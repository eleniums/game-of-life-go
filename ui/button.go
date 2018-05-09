package ui

import (
	"fmt"
	"image/color"

	"github.com/eleniums/game-of-life-go/assets"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	ptext "github.com/faiface/pixel/text"
)

// Button can be clicked to perform an action.
type Button struct {
	position    pixel.Vec
	size        pixel.Rect
	enabled     bool
	text        *ptext.Text
	textMat     pixel.Matrix
	border      *imdraw.IMDraw
	borderColor color.Color
	borderWidth float64
	fill        *imdraw.IMDraw
	fillColor   color.Color
	action      func(button *Button)
}

// NewButton creates a new button.
func NewButton(position pixel.Vec, text string, action func(button *Button)) *Button {
	button := &Button{
		size:        pixel.R(0, 0, 200, 60),
		text:        ptext.New(pixel.ZV, assets.PixelAtlas),
		enabled:     true,
		borderWidth: 5,
		borderColor: pixel.RGB(0, 0.08, 1),
		fillColor:   pixel.RGB(0, 0.45, 1),
		action:      action,
	}

	button.SetText(text)
	button.SetPosition(position)

	return button
}

// SetText sets the text for the button.
func (b *Button) SetText(text string) {
	b.text.Clear()
	fmt.Fprintln(b.text, text)
	b.redraw()
}

// SetPosition sets the position of the button.
func (b *Button) SetPosition(position pixel.Vec) {
	b.position = position
	b.redraw()
}

// SetActive enables or disables the button.
func (b *Button) SetActive(enabled bool) {
	b.enabled = enabled
}

// Position returns the position of the button.
func (b *Button) Position() pixel.Vec {
	return b.position
}

// Size returns the size of the button.
func (b *Button) Size() pixel.Rect {
	return b.size
}

// Draw the button to the screen.
func (b *Button) Draw(t pixel.Target) {
	b.border.Draw(t)
	b.fill.Draw(t)

	if b.enabled {
		b.text.Draw(t, b.textMat)
	} else {
		b.text.DrawColorMask(t, b.textMat, pixel.RGB(1, 1, 1).Mul(pixel.Alpha(0.3)))
	}
}

// Update will perform an action if the button has been clicked.
func (b *Button) Update(win *pixelgl.Window) {
	if b.enabled && win.JustPressed(pixelgl.MouseButtonLeft) {
		mpos := win.MousePosition()

		if mpos.X < b.position.X || mpos.X > b.position.X+b.size.W() || mpos.Y < b.position.Y || mpos.Y > b.position.Y+b.size.H() {
			// do nothing
		} else if b.action != nil {
			b.action(b)
		}
	}
}

// redraw will set the properties necessary to draw the button at a different position.
func (b *Button) redraw() {
	b.border = NewRectangle(b.position, b.position.Add(pixel.V(b.size.W(), b.size.H())), b.borderColor)
	b.fill = NewRectangle(b.position.Add(pixel.V(b.borderWidth, b.borderWidth)), b.position.Add(pixel.V(b.size.W()-b.borderWidth, b.size.H()-b.borderWidth)), b.fillColor)
	b.textMat = pixel.IM.Moved(b.position.Add(pixel.V((b.size.W()-b.text.Bounds().W())/2, (b.size.H()-b.text.Bounds().H())/2+5)))
}
