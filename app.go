// SPDX-License-Identifier: Apache-2.0

package ebitenguidemo

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebitenguidemo/driver"
)

type App interface {
	Update(gui *GUI) error
	Draw(screen *ebiten.Image)
}

type appWrapper struct {
	app App
}

func (a *appWrapper) Update(gui driver.GUI) error {
	return a.app.Update(&GUI{gui: gui})
}

func (a *appWrapper) Draw(screen *ebiten.Image) {
	a.app.Draw(screen)
}
