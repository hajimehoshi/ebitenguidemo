// SPDX-License-Identifier: Apache-2.0

package main

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

func drawNinePatch(dst *ebiten.Image, src *ebiten.Image, rect image.Rectangle) {
	const (
		partSizeCorner = 4
		partSizeBody   = 8
	)

	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			var sr image.Rectangle
			dx := rect.Min.X
			dy := rect.Min.Y
			dsx := 1.0
			dsy := 1.0
			op := &ebiten.DrawImageOptions{}

			switch i {
			case 0:
				sr.Min.X = 0
				sr.Max.X = partSizeCorner
			case 1:
				sr.Min.X = partSizeCorner
				sr.Max.X = partSizeCorner + partSizeBody
				dx += partSizeCorner
				dsx = float64(rect.Dx()-2*partSizeCorner) / partSizeBody
			case 2:
				sr.Min.X = partSizeCorner + partSizeBody
				sr.Max.X = partSizeCorner + partSizeBody + partSizeCorner
				dx += rect.Dx() - partSizeCorner
			}

			switch j {
			case 0:
				sr.Min.Y = 0
				sr.Max.Y = partSizeCorner
			case 1:
				sr.Min.Y = partSizeCorner
				sr.Max.Y = partSizeCorner + partSizeBody
				dy += partSizeCorner
				dsy = float64(rect.Dy()-2*partSizeCorner) / partSizeBody
			case 2:
				sr.Min.Y = partSizeCorner + partSizeBody
				sr.Max.Y = partSizeCorner + partSizeBody + partSizeCorner
				dy += rect.Dy() - partSizeCorner
			}

			op.GeoM.Scale(dsx, dsy)
			op.GeoM.Translate(float64(dx), float64(dy))
			dst.DrawImage(src.SubImage(sr).(*ebiten.Image), op)
		}
	}
}
