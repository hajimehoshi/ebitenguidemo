// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"runtime"
	"syscall/js"
)

type label struct {
	v js.Value
	x int
	y int
}

func newLabel(x, y int, text string) *label {
	l := &label{}
	runtime.SetFinalizer(l, (*label).Dispose)

	span := js.Global().Get("document").Call("createElement", "span")
	span.Get("style").Set("position", "absolute")
	span.Get("style").Set("left", fmt.Sprintf("%dpx", x))
	span.Get("style").Set("top", fmt.Sprintf("%dpx", y))
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
