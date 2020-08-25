// SPDX-License-Identifier: Apache-2.0

package ebitenguidemo

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebitenguidemo/internal/driver"
)

var textFieldImage *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(textfield_png))
	if err != nil {
		panic(err)
	}
	textFieldImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

type TextField struct {
	d driver.TextField
}

func (t *TextField) Draw(screen *ebiten.Image) {
	drawNinePatch(screen, textFieldImage, t.d.Bounds())
}

type NumberField struct {
	d driver.NumberField
}

func (n *NumberField) Value() float64 {
	return n.d.Value()
}

func (n *NumberField) SetValue(v float64) {
	n.d.SetValue(v)
}

func (n *NumberField) SetOnChange(f func(*NumberField)) {
	n.d.SetOnChange(func(driver.NumberField) {
		f(n)
	})
}

func (n *NumberField) Draw(screen *ebiten.Image) {
	drawNinePatch(screen, textFieldImage, n.d.Bounds())
}
