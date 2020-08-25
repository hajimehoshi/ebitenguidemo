// SPDX-License-Identifier: Apache-2.0

package ebitenguidemo

import (
	"image"

	"github.com/hajimehoshi/ebitenguidemo/driver"
)

type GUI struct {
	gui driver.GUI
}

func (g *GUI) NewTextField(bounds image.Rectangle) *TextField {
	return &TextField{g.gui.NewTextField(bounds)}
}

func (g *GUI) NewNumberField(bounds image.Rectangle) *NumberField {
	return &NumberField{g.gui.NewNumberField(bounds)}
}

func (g *GUI) NewLabel(x, y int, text string) *Label {
	return &Label{g.gui.NewLabel(x, y, text)}
}

func (g *GUI) NewCheckbox(x, y int) *Checkbox {
	return &Checkbox{}
}

type Label struct {
	d driver.Label
}

type Checkbox struct {
}
