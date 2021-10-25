package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/spf13/viper"
)

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
	screenWidth := viper.GetInt("layout.screen.width")
	screenHeight := viper.GetInt("layout.screen.height")

	input := NewKBHandler()

	g := NewGame(input, screenWidth, screenHeight, viper.GetInt("rules.hand_size"))

	ebiten.SetWindowTitle(viper.GetString("name"))
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
