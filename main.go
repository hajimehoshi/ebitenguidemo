// SPDX-License-Identifier: Apache-2.0

package main

import (
	"image"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	app    App
	inited bool
	items  []item
}

type item struct {
	checkbox *Checkbox
	label    *Label
}

func (g *Game) Update(_ *ebiten.Image) error {
	if !g.inited {
		t := g.app.NewTextBox(image.Rect(16, 16, 16*21, 16+24))
		t.SetOnEnter(func(t *TextBox) {
			v := strings.TrimSpace(t.Value())
			if v == "" {
				return
			}
			x, y := 16+4, 16+24*(2+len(g.items))
			g.items = append(g.items, item{
				checkbox: g.app.NewCheckbox(x, y+4),
				label:    g.app.NewLabel(x+24, y, v),
			})
			t.SetValue("")
		})
		g.inited = true
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
