// SPDX-License-Identifier: Apache-2.0

package ebitenguidemo

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebitenguidemo/gui"
)

type App interface {
	Update(gui gui.GUI) error
	Draw(screen *ebiten.Image)
}
