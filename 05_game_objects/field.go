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
	var status []PlayerStatus
	for _, c := range df.cards {
		switch c.Effects.Status {
		case StatusOriented:
			status = append(status, StatusLost)
		case StatusEscaping:
			status = append(status, StatusCaptive)
		case StatusWorking:
			status = append(status, StatusOutOfMoney)
		case StatusRecovering:
			status = append(status, StatusSick)
		}
	}
	return status
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
