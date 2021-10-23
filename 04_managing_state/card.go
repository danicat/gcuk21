package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/spf13/viper"
)

var Cards map[string]Card

type Card struct {
	Name        string
	Description string
	Type        string
	Effects     struct {
		Status string
	}
	Rules struct {
		Status []string
	}
	Count int
	Asset string

	Image *ebiten.Image

	x, y int
}

func init() {
	f, err := os.Open("cards.json")
	if err != nil {
		log.Fatal(err)
	}

	var cards map[string]Card
	d := json.NewDecoder(f)
	err = d.Decode(&cards)
	if err != nil {
		log.Fatal(err)
	}

	Cards = make(map[string]Card)

	for k, v := range cards {
		img, _, err := ebitenutil.NewImageFromFile(v.Asset)
		if err != nil {
			log.Fatal(err)
		}
		v.Image = img
		Cards[k] = v
	}
}

func (c Card) Draw(target *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	w, _ := c.Image.Size()
	scale := float64(viper.GetInt("small_card_width")) / float64(w)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64(c.x), float64(c.y))
	op.Filter = ebiten.FilterLinear
	target.DrawImage(c.Image, &op)
}
