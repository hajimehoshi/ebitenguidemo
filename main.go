// SPDX-License-Identifier: Apache-2.0

package main

import (
	"image"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	textBox *TextBox
	items   []*Item
}

type Item struct {
	checkbox *Checkbox
	label    *Label
}

func (g *Game) Update(_ *ebiten.Image) error {
	if g.textBox == nil {
		g.textBox = NewTextBox(image.Rect(16, 16, 16*21, 16+24))
		g.textBox.SetOnEnter(func(t *TextBox) {
			v := strings.TrimSpace(t.Value())
			if v == "" {
				return
			}
			x, y := 16+4, 16+24*(2+len(g.items))
			g.items = append(g.items, &Item{
				checkbox: NewCheckbox(x, y+4),
				label:    NewLabel(x+24, y, v),
			})
			t.SetValue("")
		})
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xc0, 0xc0, 0xc0, 0xff})
	g.textBox.Draw(screen)
	for _, i := range g.items {
		i.checkbox.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetRunnableOnUnfocused(true)
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
