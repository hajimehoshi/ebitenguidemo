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

var (
	checkboxOffImage *ebiten.Image
	checkboxOnImage  *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(checkboxoff_png))
	if err != nil {
		panic(err)
	}
	checkboxOffImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(checkboxon_png))
	if err != nil {
		panic(err)
	}
	checkboxOnImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

type checkbox struct {
	v js.Value
	x int
	y int

	onchange func(Checkbox)

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

func (c *checkbox) Draw(screen *ebiten.Image) {
	src := checkboxOffImage
	if c.v.Get("checked").Bool() {
		src = checkboxOnImage
	}
	drawNinePatch(screen, src, image.Rect(c.x, c.y, c.x+16, c.y+16))
}

func (c *checkbox) Checked() bool {
	return c.v.Get("checked").Bool()
}

func (c *checkbox) SetOnChange(f func(checkbox Checkbox)) {
	c.onchange = f
}
