// SPDX-License-Identifier: Apache-2.0

package js

import (
	"fmt"
	"image"
	"runtime"
	"strconv"
	"strings"
	"syscall/js"

	"github.com/hajimehoshi/ebitenguidemo/driver"
)

var isSafari bool

func init() {
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Browser_detection_using_the_user_agent
	ua := js.Global().Get("navigator").Get("userAgent").String()
	isSafari = strings.Contains(ua, "Safari/") && !strings.Contains(ua, "Chrome/") && !strings.Contains(ua, "Chromium/")
}

type textField struct {
	v                       js.Value
	bounds                  image.Rectangle
	justAfterCompositionEnd bool

	onchange func(driver.TextField)
	onenter  func(driver.TextField)

	change         js.Func
	keydown        js.Func
	compositionend js.Func
}

func newTextField(bounds image.Rectangle) *textField {
	t := &textField{}
	runtime.SetFinalizer(t, (*textField).Dispose)

	input := js.Global().Get("document").Call("createElement", "input")
	input.Set("type", "text")
	input.Get("style").Set("position", "absolute")

	body := js.Global().Get("document").Get("body")
	body.Call("appendChild", input)

	t.v = input
	t.setBounds(bounds)

	t.change = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if t.onchange != nil {
			t.onchange(t)
		}
		return nil
	})
	input.Call("addEventListener", "change", t.change)

	t.keydown = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if t.onenter == nil {
			return nil
		}

		// On Safari, keydown is fired after compositionend is fired.
		// On the other browsers, compositionend is fired after keydown is fired.
		v := t.justAfterCompositionEnd
		t.justAfterCompositionEnd = false
		if v && isSafari {
			return nil
		}

		e := args[0]
		if e.Get("isComposing").Bool() {
			return nil
		}
		if e.Get("code").String() != "Enter" {
			return nil
		}
		t.onenter(t)
		return nil
	})
	input.Call("addEventListener", "keydown", t.keydown)

	t.compositionend = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		t.justAfterCompositionEnd = true
		return nil
	})
	input.Call("addEventListener", "compositionend", t.compositionend)

	return t
}

func (t *textField) Dispose() {
	runtime.SetFinalizer(t, nil)

	body := js.Global().Get("document").Get("body")
	body.Call("removeChild", t.v)
	t.v = js.Value{}
	t.keydown.Release()
	t.compositionend.Release()
}

func (t *textField) Bounds() image.Rectangle {
	return t.bounds
}

func (t *textField) setBounds(bounds image.Rectangle) {
	t.bounds = bounds

	x := bounds.Min.X
	y := bounds.Min.Y
	w := bounds.Dx()
	h := bounds.Dy()
	t.v.Get("style").Set("left", fmt.Sprintf("%dpx", x+8))
	t.v.Get("style").Set("top", fmt.Sprintf("%dpx", y+4))
	t.v.Get("style").Set("width", fmt.Sprintf("%dpx", w-16))
	t.v.Get("style").Set("height", fmt.Sprintf("%dpx", h-8))
}

func (t *textField) Value() string {
	return t.v.Get("value").String()
}

func (t *textField) SetValue(value string) {
	t.v.Set("value", value)
}

func (t *textField) SetOnChange(f func(driver.TextField)) {
	t.onchange = f
}

func (t *textField) SetOnEnter(f func(driver.TextField)) {
	t.onenter = f
}

type numberField struct {
	*textField

	onchange func(driver.TextField)

	value float64
}

func newNumberField(bounds image.Rectangle) *numberField {
	n := &numberField{
		textField: newTextField(bounds),
	}
	runtime.SetFinalizer(n, (*numberField).Dispose)

	n.textField.v.Set("type", "number")
	n.textField.v.Set("value", "0")

	return n
}

func (n *numberField) Dispose() {
	runtime.SetFinalizer(n, nil)

	n.textField.Dispose()
	n.textField = nil
}

func (n *numberField) Bounds() image.Rectangle {
	return n.textField.Bounds()
}

func (n *numberField) Value() float64 {
	return n.value
}

func (n *numberField) SetValue(v float64) {
	changed := n.value != v
	n.value = v
	if changed {
		n.textField.v.Set("value", v)
	}
}

func (n *numberField) SetOnChange(f func(driver.NumberField)) {
	n.textField.SetOnChange(func(driver.TextField) {
		str := n.textField.v.Get("value").String()
		v, err := strconv.ParseFloat(str, 64)
		if err != nil {
			v = 0
		}
		n.value = v
		f(n)
	})
}
