package main

import "github.com/hajimehoshi/ebiten/v2"

type Field struct {
	cards []Card
}

func (f *Field) Put(card Card) {
	f.cards = append(f.cards, card)
}

func (f *Field) DrawImage(target *ebiten.Image) {
}

type DefenseField struct {
	Field
}

func NewDefenseField() *DefenseField {
	return &DefenseField{}
}

func (df *DefenseField) Resistance() []PlayerStatus {
	var res []PlayerStatus
	for _, c := range df.cards {
		res = append(res, c.Effects.Resistance)
	}
	return res
}

type TravelField struct {
	Field
}

func NewTravelField() *TravelField {
	return &TravelField{}
}

func (tf *TravelField) Distance() int {
	var sum int
	for _, c := range tf.cards {
		sum += c.Effects.Distance
	}
	return sum
}
