package main

import (
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	img *ebiten.Image
}

func (g *Game) Update() error {
	if g.img == nil {
		img, _, err := ebitenutil.NewImageFromFile("01_orientation.png")
		if err != nil {
			return err
		}
		g.img = img
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(.2, .2)
	screen.DrawImage(g.img, &op)
}

func (g *Game) Layout(width, height int) (int, int) {
	return width, height
}

func main() {
	ebiten.SetWindowTitle("Around the World")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
