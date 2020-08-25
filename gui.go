// SPDX-License-Identifier: Apache-2.0

package ebitenguidemo

import (
	"image"

	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebitenguidemo/internal/driver"
)

type drawer interface {
	Draw(screen *ebiten.Image)
}

type GUI struct {
	gui     driver.GUI
	drawers []drawer
}

func (g *GUI) NewTextField(bounds image.Rectangle) *TextField {
	t := &TextField{g.gui.NewTextField(bounds)}
	// TODO: Provide the way to remove the text field (and the other components)
	g.drawers = append(g.drawers, t)
	return t
}

func (g *GUI) NewNumberField(bounds image.Rectangle) *NumberField {
	n := &NumberField{g.gui.NewNumberField(bounds)}
	g.drawers = append(g.drawers, n)
	return n
}

func (g *GUI) NewLabel(x, y int, text string) *Label {
	return &Label{g.gui.NewLabel(x, y, text)}
}

func (g *GUI) NewCheckbox(x, y int) *Checkbox {
	c := &Checkbox{}
	g.drawers = append(g.drawers, c)
	return c
}

func (g *GUI) draw(screen *ebiten.Image) {
	for _, d := range g.drawers {
		d.Draw(screen)
	}
}

type Label struct {
	d driver.Label
}
