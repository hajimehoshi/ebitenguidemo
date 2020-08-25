// SPDX-License-Identifier: Apache-2.0

package js

import (
	"fmt"
	"image/color"

	"runtime"
	"syscall/js"
)

type label struct {
	v js.Value
	x int
	y int

	color string
}

func newLabel(x, y int, text string) *label {
	l := &label{
		color: "#000000ff",
	}
	runtime.SetFinalizer(l, (*label).Dispose)

	span := js.Global().Get("document").Call("createElement", "span")
	span.Get("style").Set("position", "absolute")
	span.Get("style").Set("left", fmt.Sprintf("%dpx", x))
	span.Get("style").Set("top", fmt.Sprintf("%dpx", y))
	span.Get("style").Set("color", l.color)
	span.Set("textContent", text)

	body := js.Global().Get("document").Get("body")
	body.Call("appendChild", span)

	l.v = span

	return l
}

func (l *label) Dispose() {
	runtime.SetFinalizer(l, nil)

	body := js.Global().Get("document").Get("body")
	body.Call("removeChild", l.v)
	l.v = js.Value{}
}

func (l *label) SetColor(clr color.Color) {
	r, g, b, a := clr.RGBA()
	rf := float64(r) / float64(a)
	gf := float64(g) / float64(a)
	bf := float64(b) / float64(a)
	af := float64(a) / 0xffff

	code := fmt.Sprintf("#%02x%02x%02x%02x", byte(rf*0xff), byte(gf*0xff), byte(bf*0xff), byte(af*0xff))
	if l.color != code {
		l.color = code
		l.v.Get("style").Set("color", code)
	}
}
