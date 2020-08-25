// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"image"
	"image/color"
)

type GUI interface {
	NewTextField(bounds image.Rectangle) TextField
	NewNumberField(bounds image.Rectangle) NumberField
	NewLabel(x, y int, text string) Label
	NewCheckbox(x, y int) Checkbox
}

type TextField interface {
	Bounds() image.Rectangle

	Value() string
	SetValue(value string)

	SetOnChange(func(textBox TextField))
	SetOnEnter(func(textBox TextField))
}

type NumberField interface {
	Bounds() image.Rectangle

	Value() float64
	SetValue(value float64)

	SetOnChange(func(numberTextField NumberField))
}

type Label interface {
	SetColor(clr color.Color)
}

type Checkbox interface {
	Bounds() image.Rectangle

	Checked() bool

	SetOnChange(func(checkbox Checkbox))
}
