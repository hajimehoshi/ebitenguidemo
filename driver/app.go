// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type App interface {
	Update(gui GUI) error
	Draw(screen *ebiten.Image)
}

type GUI interface {
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
