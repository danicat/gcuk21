package main

import (
	"errors"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/spf13/viper"

	"github.com/hajimehoshi/ebiten/v2"
)

type Deck struct {
	cards    []Card
	cardBack *ebiten.Image
}

func NewDeck() *Deck {
	d := Deck{}

	img, _, err := ebitenutil.NewImageFromFile("assets/card.backgc.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	d.cardBack = img

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

func (d *Deck) DrawImage(target *ebiten.Image) {
	if len(d.cards) > 0 {
		op := ResizeTo(d.cardBack, nil, viper.GetInt("layout.card.medium.width"), viper.GetInt("layout.card.medium.height"))
		op.GeoM.Translate(float64(viper.GetInt("layout.deck.startX")), float64(viper.GetInt("layout.deck.startY")))
		target.DrawImage(d.cardBack, op)
	}
}
