package main

import (
	"errors"
	"math/rand"
)

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	d := Deck{}
	for _, c := range Cards {
		for i := 0; i < c.Count; i++ {
			d.Insert(c)
		}
	}
	return &d
}

func (d *Deck) Insert(c Card) {
	d.cards = append(d.cards, c)
}

func (d *Deck) DrawCard() (Card, error) {
	if len(d.cards) == 0 {
		return Card{}, errors.New("deck: no cards left")
	}
	card := d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return card, nil
}

func (d *Deck) Shuffle(seed int64) {
	rand.Seed(seed)
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}
