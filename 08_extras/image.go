package main

import "github.com/hajimehoshi/ebiten/v2"

func ResizeTo(image *ebiten.Image, op *ebiten.DrawImageOptions, width, height int) *ebiten.DrawImageOptions {
	if op == nil {
		op = &ebiten.DrawImageOptions{}
	}

	w, h := image.Size()

	scaleW := float64(width) / float64(w)
	scaleH := float64(height) / float64(h)

	op.GeoM.Scale(scaleW, scaleH)
	op.Filter = ebiten.FilterLinear
	return op
}
