// SPDX-License-Identifier: Apache-2.0

package ebitenguidemo

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebitenguidemo/driver"
)

type App interface {
	Update(gui driver.GUI) error
	Draw(screen *ebiten.Image)
}
