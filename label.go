// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"runtime"
	"syscall/js"
)

type Label struct {
	v js.Value
	x int
	y int
}

func NewLabel(x, y int, text string) *Label {
	l := &Label{}
	runtime.SetFinalizer(l, (*Label).Dispose)

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

func (l *Label) Dispose() {
	runtime.SetFinalizer(l, nil)

	body := js.Global().Get("document").Get("body")
	body.Call("removeChild", l.v)
	l.v = js.Value{}
}
