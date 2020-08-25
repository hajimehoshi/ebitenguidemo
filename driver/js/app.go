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

type App struct {
	drawers []drawer
}

func (a *App) Update(screen *ebiten.Image) error {
	// Do nothing so far.
	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	for _, d := range a.drawers {
		d.Draw(screen)
	}
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (a *App) NewTextBox(bounds image.Rectangle) driver.TextBox {
	t := newTextBox(bounds)
	// TODO: How to remove the reference when t is disposed?
	a.drawers = append(a.drawers, t)
	return t
}

func (a *App) NewLabel(x, y int, text string) driver.Label {
	l := newLabel(x, y, text)
	return l
}

func (a *App) NewCheckbox(x, y int) driver.Checkbox {
	c := newCheckbox(x, y)
	a.drawers = append(a.drawers, c)
	return c
}
