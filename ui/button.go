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

func (b *Button) SetText(text string) {
	b.text.Clear()
	fmt.Fprintln(b.text, text)
	b.redraw()
}

func (b *Button) SetPosition(position pixel.Vec) {
	b.position = position
	b.redraw()
}

func (b *Button) SetActive(enabled bool) {
	b.enabled = enabled
}

func (b *Button) Draw(t pixel.Target) {
	b.border.Draw(t)
	b.fill.Draw(t)

	if b.enabled {
		b.text.Draw(t, b.textMat)
	} else {
		b.text.DrawColorMask(t, b.textMat, pixel.RGB(1, 1, 1).Mul(pixel.Alpha(0.3)))
	}
}

func (b *Button) Update(win *pixelgl.Window) {
	if b.enabled && win.JustPressed(pixelgl.MouseButtonLeft) {
		mpos := win.MousePosition()

		if mpos.X < b.position.X || mpos.X > b.position.X+b.size.W() || mpos.Y < b.position.Y || mpos.Y > b.position.Y+b.size.H() {
			// do nothing
		} else {
			b.action(b)
		}
	}
}

func (b *Button) redraw() {
	b.border = NewRectangle(b.position, b.position.Add(pixel.V(b.size.W(), b.size.H())), b.borderColor)
	b.fill = NewRectangle(b.position.Add(pixel.V(b.borderWidth, b.borderWidth)), b.position.Add(pixel.V(b.size.W()-b.borderWidth, b.size.H()-b.borderWidth)), b.fillColor)
	b.textMat = pixel.IM.Moved(b.position.Add(pixel.V((b.size.W()-b.text.Bounds().W())/2, (b.size.H()-b.text.Bounds().H())/2+5)))
}
