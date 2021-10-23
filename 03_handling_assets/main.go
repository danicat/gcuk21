package main

import (
	"log"

	_ "image/jpeg"
	_ "image/png"

	"github.com/fsnotify/fsnotify"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/spf13/viper"
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
	screen.Fill(Colors["khaki"])
	c := Cards["orientation"]
	w, h := screen.Size()
	c.x, c.y = (w-viper.GetInt("small_card_width"))/2, (h-viper.GetInt("small_card_height"))/2
	c.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	return viper.GetInt("screen_width"), viper.GetInt("screen_height")
}

func main() {
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

	ebiten.SetWindowTitle(viper.GetString("name"))
	ebiten.SetWindowSize(viper.GetInt("screen_width"), viper.GetInt("screen_height"))
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
