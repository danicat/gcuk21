package main

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/viper"
)

type Hand struct {
	cards []Card
}

func NewHand() *Hand {
	return &Hand{}
}

func (h *Hand) Put(card Card) {
	h.cards = append(h.cards, card)
}

func (h *Hand) Discard() (Card, error) {
	if len(h.cards) == 0 {
		return Card{}, errors.New("hand: no cards left to discard")
	}

	card := h.cards[len(h.cards)-1]
	h.cards = h.cards[:len(h.cards)-1]
	return card, nil
}

func (h *Hand) Left() {
	if len(h.cards) > 1 {
		first := h.cards[0]
		h.cards = append(h.cards[1:], first)
	}
}

func (h *Hand) Right() {
	if len(h.cards) > 1 {
		last := h.cards[len(h.cards)-1]
		h.cards = append([]Card{last}, h.cards[:len(h.cards)-1]...)
	}
}

func (h *Hand) DrawImage(target *ebiten.Image) {
	if len(h.cards) > 0 {
		step := float64(viper.GetInt("layout.hand.step"))
		handSize := viper.GetInt("rules.hand_size")
		for i, c := range h.cards[:len(h.cards)-1] {
			idx := i + handSize - (len(h.cards) - 1)
			op := ResizeTo(c.Image, nil, viper.GetInt("layout.card.small.width"), viper.GetInt("layout.card.small.height"))
			op.GeoM.Translate(float64(viper.GetInt("layout.hand.startX"))+step*float64(idx), float64(viper.GetInt("layout.hand.startY")))
			target.DrawImage(c.Image, op)
		}

		// last card is always the big one / selected
		card := h.cards[len(h.cards)-1]
		step2 := float64(viper.GetInt("layout.hand.step2"))
		op := ResizeTo(card.Image, nil, viper.GetInt("layout.card.large.width"), viper.GetInt("layout.card.large.height"))
		op.GeoM.Translate(float64(viper.GetInt("layout.hand.startX"))+step2*float64(handSize), float64(viper.GetInt("layout.hand.startY")-viper.GetInt("layout.card.large.height")/2))
		target.DrawImage(card.Image, op)
	}
}
