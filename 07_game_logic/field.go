package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/viper"
)

type Field struct {
	cards []Card
}

func (f *Field) Put(card Card) {
	f.cards = append(f.cards, card)
}

type DefenseField struct {
	player int
	Field
}

func NewDefenseField(player int) *DefenseField {
	return &DefenseField{player: player}
}

func (df *DefenseField) Resistance() []PlayerStatus {
	var res []PlayerStatus
	for _, c := range df.cards {
		res = append(res, c.Effects.Resistance)
	}
	return res
}

func (df *DefenseField) DrawImage(target *ebiten.Image) {
	if len(df.cards) > 0 {
		step := float64(viper.GetInt(fmt.Sprintf("layout.player%d.defense.step", df.player)))
		startX := fmt.Sprintf("layout.player%d.defense.startX", df.player)
		startY := fmt.Sprintf("layout.player%d.defense.startY", df.player)
		for i, c := range df.cards {
			op := ResizeTo(c.Image, nil, viper.GetInt("layout.card.tiny.width"), viper.GetInt("layout.card.tiny.height"))
			op.GeoM.Translate(float64(viper.GetInt(startX))+step*float64(i), float64(viper.GetInt(startY)))
			target.DrawImage(c.Image, op)
		}
	}
}

type TravelField struct {
	player int
	Field
}

func NewTravelField(player int) *TravelField {
	return &TravelField{player: player}
}

func (tf *TravelField) Distance() int {
	var sum int
	for _, c := range tf.cards {
		sum += c.Effects.Distance
	}
	return sum
}

func (tf *TravelField) DrawImage(target *ebiten.Image) {
	if len(tf.cards) > 0 {
		step := float64(viper.GetInt(fmt.Sprintf("layout.player%d.travel.step", tf.player)))
		startX := fmt.Sprintf("layout.player%d.travel.startX", tf.player)
		startY := fmt.Sprintf("layout.player%d.travel.startY", tf.player)
		for i, c := range tf.cards {
			op := ResizeTo(c.Image, nil, viper.GetInt("layout.card.medium.width"), viper.GetInt("layout.card.medium.height"))
			op.GeoM.Translate(float64(viper.GetInt(startX))+step*float64(i), float64(viper.GetInt(startY)))
			target.DrawImage(c.Image, op)
		}
	}
}
