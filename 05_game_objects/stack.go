package main

import (
	"errors"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
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

func (s *Stack) DrawImage(target *ebiten.Image) {
	card, err := s.Top()
	if err != nil {
		log.Printf("stack: can't draw image - %s", err)
		return
	}

	target.DrawImage(card.Image, nil)
}

type BattleStack struct {
	Stack
}

func NewBattleStack() *BattleStack {
	return &BattleStack{}
}

func (bs *BattleStack) Status() PlayerStatus {
	c, err := bs.Top()
	if err != nil {
		log.Printf("battle stack: current status unknown - %s", err)
		return StatusUnknown
	}
	return c.Effects.Status
}

type TerrainStack struct {
	Stack
}

func NewTerrainStack() *TerrainStack {
	return &TerrainStack{}
}

func (ts *TerrainStack) Terrain() string {
	c, err := ts.Top()
	if err != nil {
		log.Printf("terrain stack: current terrain unknown - %s", err)
		return "any"
	}
	return c.Effects.Terrain
}
