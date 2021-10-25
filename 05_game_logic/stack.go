package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/viper"
)

type Stack struct {
	cards []Card
}

func (s *Stack) Top() (Card, error) {
	if len(s.cards) == 0 {
		return Card{}, errors.New("stack: empty stack")
	}

	return s.cards[len(s.cards)-1], nil
}

func (s *Stack) Put(card Card) {
	s.cards = append(s.cards, card)
}

func (s *Stack) Pop() (Card, error) {
	if len(s.cards) == 0 {
		return Card{}, errors.New("stack: empty stack")
	}
	card := s.cards[len(s.cards)-1]
	s.cards = s.cards[:len(s.cards)-1]
	return card, nil
}

type BattleStack struct {
	player int
	Stack
}

func NewBattleStack(player int) *BattleStack {
	return &BattleStack{player: player}
}

func (bs *BattleStack) Status() PlayerStatus {
	c, err := bs.Top()
	if err != nil {
		log.Printf("battle stack: current status unknown - %s", err)
		return StatusUnknown
	}
	return c.Effects.Status
}

func (bs *BattleStack) DrawImage(target *ebiten.Image) {
	card, err := bs.Top()
	if err != nil {
		// log.Printf("battle stack: can't draw image - %s", err)
		return
	}

	startX := fmt.Sprintf("layout.player%d.battle.startX", bs.player)
	startY := fmt.Sprintf("layout.player%d.battle.startY", bs.player)

	op := ResizeTo(card.Image, nil, viper.GetInt("layout.card.medium.width"), viper.GetInt("layout.card.medium.height"))
	op.GeoM.Translate(float64(viper.GetInt(startX)), float64(viper.GetInt(startY)))
	target.DrawImage(card.Image, op)
}

type TerrainStack struct {
	player int
	Stack
}

func NewTerrainStack(player int) *TerrainStack {
	return &TerrainStack{player: player}
}

func (ts *TerrainStack) Terrain() string {
	c, err := ts.Top()
	if err != nil {
		log.Printf("terrain stack: current terrain unknown - %s", err)
		return "any"
	}
	return c.Effects.Terrain
}

func (ts *TerrainStack) DrawImage(target *ebiten.Image) {
	card, err := ts.Top()
	if err != nil {
		// log.Printf("terrain stack: can't draw image - %s", err)
		return
	}

	startX := fmt.Sprintf("layout.player%d.terrain.startX", ts.player)
	startY := fmt.Sprintf("layout.player%d.terrain.startY", ts.player)

	op := ResizeTo(card.Image, nil, viper.GetInt("layout.card.medium.width"), viper.GetInt("layout.card.medium.height"))
	op.GeoM.Translate(float64(viper.GetInt(startX)), float64(viper.GetInt(startY)))
	target.DrawImage(card.Image, op)
}

type Graveyard struct {
	Stack
}

func NewGraveyard() *Graveyard {
	return &Graveyard{}
}

func (g *Graveyard) DrawImage(target *ebiten.Image) {
	card, err := g.Top()
	if err != nil {
		return
	}

	op := ResizeTo(card.Image, nil, viper.GetInt("layout.card.medium.width"), viper.GetInt("layout.card.medium.height"))
	op.GeoM.Translate(float64(viper.GetInt("layout.graveyard.startX")), float64(viper.GetInt("layout.graveyard.startY")))
	target.DrawImage(card.Image, op)
}
