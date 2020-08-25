// SPDX-License-Identifier: Apache-2.0

package gui

import (
	"image"
	"image/color"
)

type GUI interface {
	NewTextBox(bounds image.Rectangle) TextBox
	NewNumberTextBox(bounds image.Rectangle) NumberTextBox
	NewLabel(x, y int, text string) Label
	NewCheckbox(x, y int) Checkbox
}

type TextBox interface {
	Value() string
	SetValue(value string)

	SetOnChange(func(textBox TextBox))
	SetOnEnter(func(textBox TextBox))
}

type NumberTextBox interface {
	Value() float64
	SetValue(value float64)

	SetOnChange(func(numberTextBox NumberTextBox))
}

type Label interface {
	SetColor(clr color.Color)
}

type Checkbox interface {
	Checked() bool

	SetOnChange(func(checkbox Checkbox))
}
