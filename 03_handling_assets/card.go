package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/viper"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var Cards map[string]Card

type Card struct {
	Name    string
	Type    string
	Effects struct {
		Status string
	}
	Rules struct {
		Status []string
	}
	Count int
	Asset string

	Image *ebiten.Image
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

func (c *Card) Draw(target *ebiten.Image) {
	op := ResizeTo(c.Image, nil, viper.GetInt("card.width"), viper.GetInt("card.height"))
	op.GeoM.Translate(float64(viper.GetInt("card.startX")), float64(viper.GetInt("card.startY")))
	target.DrawImage(c.Image, op)
}
