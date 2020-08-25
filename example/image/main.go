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

	"github.com/hajimehoshi/ebitenguidemo/driver"
	"github.com/hajimehoshi/ebitenguidemo/driver/js"
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
	labels    []driver.Label
	textBoxes []driver.NumberTextBox
}

func (a *App) initIfNeeded(gui driver.GUI) {
	if len(a.textBoxes) > 0 {
		return
	}

	const unit = 24
	a.labels = []driver.Label{
		gui.NewLabel(unit, unit, "A"),
		gui.NewLabel(unit*6, unit, "B"),
		gui.NewLabel(unit*11, unit, "TX"),
		gui.NewLabel(unit, unit*4, "C"),
		gui.NewLabel(unit*6, unit*4, "D"),
		gui.NewLabel(unit*11, unit*4, "TY"),
	}
	// TODO: Ideally the text box's text head should be on the same line as the label's text.
	// Adjust the position.
	a.textBoxes = []driver.NumberTextBox{
		gui.NewNumberTextBox(image.Rect(unit, unit*2, unit*(1+4), unit*(2+1))),     // a
		gui.NewNumberTextBox(image.Rect(unit*6, unit*2, unit*(6+4), unit*(2+1))),   // b
		gui.NewNumberTextBox(image.Rect(unit*11, unit*2, unit*(11+4), unit*(2+1))), // tx
		gui.NewNumberTextBox(image.Rect(unit, unit*5, unit*(1+4), unit*(5+1))),     // c
		gui.NewNumberTextBox(image.Rect(unit*6, unit*5, unit*(6+4), unit*(5+1))),   // d
		gui.NewNumberTextBox(image.Rect(unit*11, unit*5, unit*(11+4), unit*(5+1))), // ty
	}
	a.textBoxes[0].SetValue(1)
	a.textBoxes[1].SetValue(0)
	a.textBoxes[2].SetValue(0)
	a.textBoxes[3].SetValue(0)
	a.textBoxes[4].SetValue(1)
	a.textBoxes[5].SetValue(0)

	for i, t := range a.textBoxes {
		i := i
		t.SetOnChange(func(n driver.NumberTextBox) {
			a.geoM.SetElement(i/3, i%3, n.Value())
		})
	}
}

func (a *App) Update(gui driver.GUI) error {
	a.initIfNeeded(gui)
	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xc0, 0xc0, 0xc0, 0xff})

	op := &ebiten.DrawImageOptions{}
	op.GeoM = a.geoM
	screen.DrawImage(gophersImage, op)
}

func main() {
	ebiten.SetRunnableOnUnfocused(true)
	if err := ebiten.RunGame(js.NewApp(&App{})); err != nil {
		panic(err)
	}
}
