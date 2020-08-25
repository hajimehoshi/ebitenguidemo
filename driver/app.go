// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type App interface {
	ebiten.Game
	Draw(screen *ebiten.Image) // Draw is not defined ebiten.Game so far, and will be defined in Ebiten v2.

	NewTextBox(bounds image.Rectangle) TextBox
	NewLabel(x, y int, text string) Label
	NewCheckbox(x, y int) Checkbox
}

type TextBox interface {
	Value() string
	SetValue(value string)

	// TODO: Should this be SetOnChange?
	SetOnEnter(func(textBox TextBox))
}

type Label interface {
	SetColor(clr color.Color)
}

type Checkbox interface {
	Checked() bool

	SetOnChange(func(checkbox Checkbox))
}
