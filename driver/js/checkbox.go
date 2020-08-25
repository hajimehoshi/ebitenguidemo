// SPDX-License-Identifier: Apache-2.0

package js

import (
	"fmt"
	"image"
	"runtime"
	"syscall/js"

	"github.com/hajimehoshi/ebitenguidemo/driver"
)

type checkbox struct {
	v       js.Value
	x       int
	y       int
	checked bool // TODO: Consider 'intermediate' state. (aria-checked='mixed')

	onchange func(driver.Checkbox)

	change js.Func
}

func newCheckbox(x, y int) *checkbox {
	c := &checkbox{
		x: x,
		y: y,
	}
	runtime.SetFinalizer(c, (*checkbox).Dispose)

	input := js.Global().Get("document").Call("createElement", "input")
	input.Set("type", "checkbox")
	input.Get("style").Set("position", "absolute")
	input.Get("style").Set("left", fmt.Sprintf("%dpx", x))
	input.Get("style").Set("top", fmt.Sprintf("%dpx", y))
	input.Get("style").Set("opacity", "0")

	body := js.Global().Get("document").Get("body")
	body.Call("appendChild", input)

	c.v = input

	c.change = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		c.checked = c.v.Get("checked").Bool()
		if c.onchange != nil {
			c.onchange(c)
		}
		return nil
	})
	c.v.Call("addEventListener", "change", c.change)

	return c
}

func (c *checkbox) Dispose() {
	runtime.SetFinalizer(c, nil)

	body := js.Global().Get("document").Get("body")
	body.Call("removeChild", c.v)
	c.v = js.Value{}
	c.change.Release()
}

func (c *checkbox) Bounds() image.Rectangle {
	return image.Rect(c.x, c.y, c.x +16, c.y+16)
}

func (c *checkbox) Checked() bool {
	return c.checked
}

func (c *checkbox) SetOnChange(f func(checkbox driver.Checkbox)) {
	c.onchange = f
}
