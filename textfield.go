// SPDX-License-Identifier: Apache-2.0

package ebitenguidemo

import (
	"github.com/hajimehoshi/ebitenguidemo/driver"
)

type TextField struct {
	d driver.TextField
}

type NumberField struct {
	d driver.NumberField
}

func (n *NumberField) Value() float64 {
	return n.d.Value()
}

func (n *NumberField) SetValue(v float64) {
	n.d.SetValue(v)
}

func (n *NumberField) SetOnChange(f func(*NumberField)) {
	n.d.SetOnChange(func(driver.NumberField) {
		f(n)
	})
}
