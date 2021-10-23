package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/spf13/viper"
)

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

	g := NewGame()

	ebiten.SetWindowTitle(viper.GetString("name"))
	ebiten.SetWindowSize(viper.GetInt("screen_width"), viper.GetInt("screen_height"))
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
