// SPDX-License-Identifier: Apache-2.0

package js

import (
	"image"

	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebitenguidemo/driver"
)

type drawer interface {
	Draw(screen *ebiten.Image)
}

type guiImpl struct {
	drawers []drawer
}

func (g *guiImpl) NewTextField(bounds image.Rectangle) driver.TextField {
	t := newTextField(bounds)
	// TODO: How to remove the reference when t is disposed?
	g.drawers = append(g.drawers, t)
	return t
}

func (g *guiImpl) NewNumberField(bounds image.Rectangle) driver.NumberField {
	n := newNumberField(bounds)
	g.drawers = append(g.drawers, n)
	return n
}

func (g *guiImpl) NewLabel(x, y int, text string) driver.Label {
	l := newLabel(x, y, text)
	return l
}

func (g *guiImpl) NewCheckbox(x, y int) driver.Checkbox {
	c := newCheckbox(x, y)
	g.drawers = append(g.drawers, c)
	return c
}

func (g *guiImpl) Draw(screen *ebiten.Image) {
	for _, d := range g.drawers {
		d.Draw(screen)
	}
}
