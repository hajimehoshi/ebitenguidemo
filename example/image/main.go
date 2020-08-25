// SPDX-License-Identifier: Apache-2.0

// +build example

package main

import (
	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/images"

	"github.com/hajimehoshi/ebitenguidemo"
	"github.com/hajimehoshi/ebitenguidemo/gui"
)

var (
	gophersImage *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(images.Gophers_jpg))
	if err != nil {
		panic(err)
	}
	gophersImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

type App struct {
	// Model
	geoM ebiten.GeoM

	// View
	labels    []gui.Label
	textBoxes []gui.NumberTextBox
}

func (a *App) initIfNeeded(g gui.GUI) {
	if len(a.textBoxes) > 0 {
		return
	}

	const unit = 24
	a.labels = []gui.Label{
		g.NewLabel(unit, unit, "A"),
		g.NewLabel(unit*6, unit, "B"),
		g.NewLabel(unit*11, unit, "TX"),
		g.NewLabel(unit, unit*4, "C"),
		g.NewLabel(unit*6, unit*4, "D"),
		g.NewLabel(unit*11, unit*4, "TY"),
	}
	// TODO: Ideally the text box's text head should be on the same line as the label's text.
	// Adjust the position.
	a.textBoxes = []gui.NumberTextBox{
		g.NewNumberTextBox(image.Rect(unit, unit*2, unit*(1+4), unit*(2+1))),     // a
		g.NewNumberTextBox(image.Rect(unit*6, unit*2, unit*(6+4), unit*(2+1))),   // b
		g.NewNumberTextBox(image.Rect(unit*11, unit*2, unit*(11+4), unit*(2+1))), // tx
		g.NewNumberTextBox(image.Rect(unit, unit*5, unit*(1+4), unit*(5+1))),     // c
		g.NewNumberTextBox(image.Rect(unit*6, unit*5, unit*(6+4), unit*(5+1))),   // d
		g.NewNumberTextBox(image.Rect(unit*11, unit*5, unit*(11+4), unit*(5+1))), // ty
	}
	a.textBoxes[0].SetValue(1)
	a.textBoxes[1].SetValue(0)
	a.textBoxes[2].SetValue(0)
	a.textBoxes[3].SetValue(0)
	a.textBoxes[4].SetValue(1)
	a.textBoxes[5].SetValue(0)

	for i, t := range a.textBoxes {
		i := i
		t.SetOnChange(func(n gui.NumberTextBox) {
			a.geoM.SetElement(i/3, i%3, n.Value())
		})
	}
}

func (a *App) Update(g gui.GUI) error {
	a.initIfNeeded(g)
	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xc0, 0xc0, 0xc0, 0xff})

	op := &ebiten.DrawImageOptions{}
	op.GeoM = a.geoM
	screen.DrawImage(gophersImage, op)
}

func main() {	
	if err := ebitenguidemo.Run(&App{}); err != nil {
		panic(err)
	}
}
