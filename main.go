// SPDX-License-Identifier: Apache-2.0

package main

import (
	"image"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebitenguidemo/driver"
	"github.com/hajimehoshi/ebitenguidemo/driver/js"
)

type App struct {
	// Model
	items []*item

	// View
	textBox   driver.TextBox
	itemViews []*itemView
}

type item struct {
	checked bool
}

type itemView struct {
	checkbox driver.Checkbox
	label    driver.Label
}

func (a *App) initIfNeeded(gui driver.GUI) {
	if a.textBox != nil {
		return
	}

	a.textBox = gui.NewTextBox(image.Rect(16, 16, 16*21, 16+24))
	a.textBox.SetOnEnter(func(t driver.TextBox) {
		v := strings.TrimSpace(t.Value())
		if v == "" {
			return
		}

		i := &item{}
		a.items = append(a.items, i)

		x, y := 16+4, 16+24*(2+len(a.itemViews))
		iv := &itemView{
			checkbox: gui.NewCheckbox(x, y+4),
			label:    gui.NewLabel(x+24, y, v),
		}
		a.itemViews = append(a.itemViews, iv)
		iv.checkbox.SetOnChange(func(c driver.Checkbox) {
			i.checked = c.Checked()
		})
		t.SetValue("")
	})
}

func (a *App) Update(gui driver.GUI) error {
	a.initIfNeeded(gui)

	// Update the view based on the model.
	for i, item := range a.items {
		clr := color.RGBA{0, 0, 0, 0xff}
		if item.checked {
			clr = color.RGBA{0x80, 0x80, 0x80, 0xff}
		}
		a.itemViews[i].label.SetColor(clr)
	}

	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xc0, 0xc0, 0xc0, 0xff})
}

func main() {
	ebiten.SetRunnableOnUnfocused(true)
	if err := ebiten.RunGame(js.NewApp(&App{})); err != nil {
		panic(err)
	}
}
