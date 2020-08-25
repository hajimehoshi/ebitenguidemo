// SPDX-License-Identifier: Apache-2.0

package js

import (
	"image"

	"github.com/hajimehoshi/ebitenguidemo/internal/driver"
)

type guiImpl struct {
}

func (g *guiImpl) NewTextField(bounds image.Rectangle) driver.TextField {
	t := newTextField(bounds)
	return t
}

func (g *guiImpl) NewNumberField(bounds image.Rectangle) driver.NumberField {
	n := newNumberField(bounds)
	return n
}

func (g *guiImpl) NewLabel(x, y int, text string) driver.Label {
	l := newLabel(x, y, text)
	return l
}

func (g *guiImpl) NewCheckbox(x, y int) driver.Checkbox {
	c := newCheckbox(x, y)
	return c
}
