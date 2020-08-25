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

type gui struct {
	drawers []drawer
}

func (g *gui) NewTextBox(bounds image.Rectangle) driver.TextBox {
	t := newTextBox(bounds)
	// TODO: How to remove the reference when t is disposed?
	g.drawers = append(g.drawers, t)
	return t
}

func (g *gui) NewLabel(x, y int, text string) driver.Label {
	l := newLabel(x, y, text)
	return l
}

func (g *gui) NewCheckbox(x, y int) driver.Checkbox {
	c := newCheckbox(x, y)
	g.drawers = append(g.drawers, c)
	return c
}

func (g *gui) Draw(screen *ebiten.Image) {
	for _, d := range g.drawers {
		d.Draw(screen)
	}
}
