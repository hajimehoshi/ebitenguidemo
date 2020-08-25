// SPDX-License-Identifier: Apache-2.0

package js

import (
	"image"

	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebitenguidemo/gui"
)

type drawer interface {
	Draw(screen *ebiten.Image)
}

type guiImpl struct {
	drawers []drawer
}

func (g *guiImpl) NewTextBox(bounds image.Rectangle) gui.TextBox {
	t := newTextBox(bounds)
	// TODO: How to remove the reference when t is disposed?
	g.drawers = append(g.drawers, t)
	return t
}

func (g *guiImpl) NewNumberTextBox(bounds image.Rectangle) gui.NumberTextBox {
	n := newNumberTextBox(bounds)
	// TODO: How to remove the reference when t is disposed?
	g.drawers = append(g.drawers, n)
	return n
}

func (g *guiImpl) NewLabel(x, y int, text string) gui.Label {
	l := newLabel(x, y, text)
	return l
}

func (g *guiImpl) NewCheckbox(x, y int) gui.Checkbox {
	c := newCheckbox(x, y)
	g.drawers = append(g.drawers, c)
	return c
}

func (g *guiImpl) Draw(screen *ebiten.Image) {
	for _, d := range g.drawers {
		d.Draw(screen)
	}
}
