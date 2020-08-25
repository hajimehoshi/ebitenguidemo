// SPDX-License-Identifier: Apache-2.0

package ebitenguidemo

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebitenguidemo/driver"
)

var (
	checkboxOffImage *ebiten.Image
	checkboxOnImage  *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(checkboxoff_png))
	if err != nil {
		panic(err)
	}
	checkboxOffImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(checkboxon_png))
	if err != nil {
		panic(err)
	}
	checkboxOnImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

type Checkbox struct {
	d driver.Checkbox
}

func (c *Checkbox) Draw(screen *ebiten.Image) {
	src := checkboxOffImage
	if c.d.Checked() {
		src = checkboxOnImage
	}
	drawNinePatch(screen, src, c.d.Bounds())
}
