// SPDX-License-Identifier: Apache-2.0

package js

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"runtime"
	"strings"
	"syscall/js"

	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebitenguidemo/driver"
)

var textBoxImage *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(textbox_png))
	if err != nil {
		panic(err)
	}
	textBoxImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

var isSafari bool

func init() {
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Browser_detection_using_the_user_agent
	ua := js.Global().Get("navigator").Get("userAgent").String()
	isSafari = strings.Contains(ua, "Safari/") && !strings.Contains(ua, "Chrome/") && !strings.Contains(ua, "Chromium/")
}

type textBox struct {
	v                       js.Value
	bounds                  image.Rectangle
	justAfterCompositionEnd bool

	onenter func(driver.TextBox)

	keydown        js.Func
	compositionend js.Func
}

func newTextBox(bounds image.Rectangle) *textBox {
	t := &textBox{}
	runtime.SetFinalizer(t, (*textBox).Dispose)

	input := js.Global().Get("document").Call("createElement", "input")
	input.Set("type", "text")
	input.Get("style").Set("position", "absolute")

	body := js.Global().Get("document").Get("body")
	body.Call("appendChild", input)

	t.v = input
	t.setBounds(bounds)

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

func (t *textBox) Dispose() {
	runtime.SetFinalizer(t, nil)

	body := js.Global().Get("document").Get("body")
	body.Call("removeChild", t.v)
	t.v = js.Value{}
	t.keydown.Release()
	t.compositionend.Release()
}

func (t *textBox) setBounds(bounds image.Rectangle) {
	t.bounds = bounds

	x := bounds.Min.X
	y := bounds.Min.Y
	w := bounds.Dx()
	h := bounds.Dy()
	t.v.Get("style").Set("left", fmt.Sprintf("%dpx", y+8))
	t.v.Get("style").Set("top", fmt.Sprintf("%dpx", x+4))
	t.v.Get("style").Set("width", fmt.Sprintf("%dpx", w-16))
	t.v.Get("style").Set("height", fmt.Sprintf("%dpx", h-8))
}

func (t *textBox) Draw(screen *ebiten.Image) {
	drawNinePatch(screen, textBoxImage, t.bounds)
}

func (t *textBox) Value() string {
	return t.v.Get("value").String()
}

func (t *textBox) SetValue(value string) {
	t.v.Set("value", value)
}

func (t *textBox) SetOnEnter(f func(driver.TextBox)) {
	t.onenter = f
}
