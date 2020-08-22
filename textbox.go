// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"runtime"
	"syscall/js"

	"github.com/hajimehoshi/ebiten"
)

var textBoxImage *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(textbox_png))
	if err != nil {
		panic(err)
	}
	textBoxImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

type TextBox struct {
	v      js.Value
	bounds image.Rectangle

	onenter func(*TextBox)

	keydown js.Func
}

func NewTextBox(bounds image.Rectangle) *TextBox {
	t := &TextBox{}
	runtime.SetFinalizer(t, (*TextBox).Dispose)

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

	return t
}

func (t *TextBox) Dispose() {
	runtime.SetFinalizer(t, nil)

	body := js.Global().Get("document").Get("body")
	body.Call("removeChild", t.v)
	t.v = js.Value{}
	t.keydown.Release()
}

func (t *TextBox) setBounds(bounds image.Rectangle) {
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

func (t *TextBox) Draw(screen *ebiten.Image) {
	drawNinePatch(screen, textBoxImage, t.bounds)
}

func (t *TextBox) Value() string {
	return t.v.Get("value").String()
}

func (t *TextBox) SetValue(value string) {
	t.v.Set("value", value)
}

func (t *TextBox) SetOnEnter(f func(*TextBox)) {
	t.onenter = f
}
