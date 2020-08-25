// SPDX-License-Identifier: Apache-2.0

package js

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebitenguidemo/driver"
)

type App struct {
	app driver.App
	gui gui
}

func NewApp(app driver.App) *App {
	return &App{
		app: app,
	}
}

func (a *App) Update(screen *ebiten.Image) error {
	if err := a.app.Update(&a.gui); err != nil {
		return err
	}
	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	a.app.Draw(screen)
	a.gui.Draw(screen)
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
