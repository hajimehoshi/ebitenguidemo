// SPDX-License-Identifier: Apache-2.0

package js

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebitenguidemo/gui"
)

type app interface {
	Update(gui gui.GUI) error
	Draw(screen *ebiten.Image)
}

type App struct {
	app app
	gui guiImpl
}

func NewApp(app app) *App {
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
