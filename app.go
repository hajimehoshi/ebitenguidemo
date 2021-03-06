// SPDX-License-Identifier: Apache-2.0

package ebitenguidemo

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebitenguidemo/internal/driver"
)

type App interface {
	Update(gui *GUI) error
	Draw(screen *ebiten.Image)
}

type appWrapper struct {
	app App
	gui GUI
}

func (a *appWrapper) Update(gui driver.GUI) error {
	a.gui.gui = gui
	return a.app.Update(&a.gui)
}

func (a *appWrapper) Draw(screen *ebiten.Image) {
	a.app.Draw(screen)
	a.gui.draw(screen)
}
