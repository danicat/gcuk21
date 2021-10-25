package main

import (
	"log"
	"sync"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	img  *ebiten.Image
	once sync.Once
}

func (g *Game) Update() error {
	g.once.Do(func() {
		img, _, err := ebitenutil.NewImageFromFile("01_orientation.png")
		if err != nil {
			log.Fatal(err)
		}
		g.img = img
	})
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(.3, .3)
	op.GeoM.Translate(300, 300)
	screen.DrawImage(g.img, &op)
}

func (g *Game) Layout(width, height int) (int, int) {
	return 1280, 720
}

func main() {
	ebiten.SetWindowTitle("Around the World")
	ebiten.SetWindowSize(1280, 720)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
