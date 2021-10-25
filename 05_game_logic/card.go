package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var Cards map[string]Card

type Card struct {
	Name    string
	Type    string
	Effects struct {
		Status     PlayerStatus
		Terrain    string
		Distance   int
		Resistance PlayerStatus
	}
	Rules struct {
		Status   []PlayerStatus
		Terrains []string
		Target   string
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
