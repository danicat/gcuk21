package main

import (
	"log"
	"sync"

	_ "image/jpeg"
	_ "image/png"

	"github.com/fsnotify/fsnotify"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/spf13/viper"
)

type Game struct {
	background *ebiten.Image
	once       sync.Once
}

func (g *Game) Update() error {
	g.once.Do(func() {
		img, _, err := ebitenutil.NewImageFromFile("background.png")
		if err != nil {
			log.Fatal(err)
		}
		g.background = img
	})

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := ResizeTo(g.background, nil, viper.GetInt("screen_width"), viper.GetInt("screen_height"))
	screen.DrawImage(g.background, op)

	c := Cards["orientation"]
	c.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	return viper.GetInt("screen_width"), viper.GetInt("screen_height")
}

func ResizeTo(image *ebiten.Image, op *ebiten.DrawImageOptions, width, height int) *ebiten.DrawImageOptions {
	if op == nil {
		op = &ebiten.DrawImageOptions{}
	}

	w, h := image.Size()

	scaleW := float64(width) / float64(w)
	scaleH := float64(height) / float64(h)

	op.GeoM.Scale(scaleW, scaleH)
	return op
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
	viper.WatchConfig()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	ebiten.SetWindowTitle(viper.GetString("name"))
	ebiten.SetWindowSize(viper.GetInt("screen_width"), viper.GetInt("screen_height"))
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
