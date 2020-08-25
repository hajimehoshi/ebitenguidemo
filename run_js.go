// SPDX-License-Identifier: Apache-2.0

package ebitenguidemo

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebitenguidemo/driver/js"
)

func Run(app App) error {
	ebiten.SetRunnableOnUnfocused(true)
	if err := ebiten.RunGame(js.NewApp(&appWrapper{app: app})); err != nil {
		return err
	}
	return nil
}
