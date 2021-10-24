package main

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

type Hand struct {
	cards    []Card
	selected int
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

func (h *Hand) Selected() (Card, error) {
	if len(h.cards) == 0 {
		return Card{}, errors.New("hand: no cards left in hand")
	}
	return h.cards[h.selected], nil
}

func (h *Hand) DrawImage(target *ebiten.Image) {

}
