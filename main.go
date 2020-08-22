// SPDX-License-Identifier: Apache-2.0

package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	textBox *TextBox
	items   []*Label
}

func (g *Game) Update(_ *ebiten.Image) error {
	if g.textBox == nil {
		g.textBox = NewTextBox(image.Rect(16, 16, 16*21, 16+24))
		g.textBox.SetOnEnter(func(t *TextBox) {
			v := t.Value()
			if v == "" {
				return
			}
			g.items = append(g.items, NewLabel(16+4, 16+24*(2+len(g.items)), v))
			t.SetValue("")
		})
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xc0, 0xc0, 0xc0, 0xff})
	g.textBox.Draw(screen)
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
