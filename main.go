// SPDX-License-Identifier: Apache-2.0

package main

import (
	"image"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	app App

	// Model
	items []*item

	// View
	textBox   TextBox
	itemViews []*itemView
}

type item struct {
	checked bool
}

type itemView struct {
	checkbox Checkbox
	label    Label
}

func (g *Game) Update(_ *ebiten.Image) error {
	if g.textBox == nil {
		g.textBox = g.app.NewTextBox(image.Rect(16, 16, 16*21, 16+24))
		g.textBox.SetOnEnter(func(t TextBox) {
			v := strings.TrimSpace(t.Value())
			if v == "" {
				return
			}

			i := &item{}
			g.items = append(g.items, i)

			x, y := 16+4, 16+24*(2+len(g.itemViews))
			iv := &itemView{
				checkbox: g.app.NewCheckbox(x, y+4),
				label:    g.app.NewLabel(x+24, y, v),
			}
			g.itemViews = append(g.itemViews, iv)
			iv.checkbox.SetOnChange(func(c Checkbox) {
				i.checked = c.Checked()
			})
			t.SetValue("")
		})
	}

	// Update the view based on the model.
	for i, item := range g.items {
		clr := color.RGBA{0, 0, 0, 0xff}
		if item.checked {
			clr = color.RGBA{0x80, 0x80, 0x80, 0xff}
		}
		g.itemViews[i].label.SetColor(clr)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xc0, 0xc0, 0xc0, 0xff})
	g.app.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.app.Layout(outsideWidth, outsideHeight)
}

func main() {
	ebiten.SetRunnableOnUnfocused(true)
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
