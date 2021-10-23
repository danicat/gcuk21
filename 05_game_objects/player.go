package main

import (
	"encoding/json"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player interface {
	Update()
	DrawImage(target *ebiten.Image)
	Hand() *Hand
	Distance() int
	Status() PlayerStatus
	Resistance() []PlayerStatus
	Terrain() string
}

type PlayerStatus int

const (
	StatusUnknown PlayerStatus = iota
	StatusOriented
	StatusWorking
	StatusEscaping
	StatusRecovering
	StatusLost
	StatusOutOfMoney
	StatusCaptive
	StatusSick
)

func (ps *PlayerStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	case "oriented":
		*ps = StatusOriented
	case "working":
		*ps = StatusWorking
	case "escaping":
		*ps = StatusEscaping
	case "recovering":
		*ps = StatusRecovering
	case "lost":
		*ps = StatusLost
	case "outofmoney":
		*ps = StatusOutOfMoney
	case "captive":
		*ps = StatusCaptive
	case "sick":
		*ps = StatusSick
	}

	return nil
}

type BasePlayer struct {
	g  *Game
	h  *Hand
	bs *BattleStack
	ts *TerrainStack
	tf *TravelField
	df *DefenseField
}

func NewBasePlayer(g *Game) *BasePlayer {
	return &BasePlayer{
		g:  g,
		h:  NewHand(),
		bs: NewBattleStack(),
		ts: NewTerrainStack(),
		tf: NewTravelField(),
		df: NewDefenseField(),
	}
}

func (bp *BasePlayer) DrawImage(target *ebiten.Image) {

}

func (bp *BasePlayer) Hand() *Hand {
	return bp.h
}

func (bp *BasePlayer) Distance() int {
	return bp.tf.Distance()
}

func (bp *BasePlayer) Status() PlayerStatus {
	return bp.bs.Status()
}

func (bp *BasePlayer) Resistance() []PlayerStatus {
	return bp.df.Resistance()
}

func (bp *BasePlayer) Terrain() string {
	return bp.ts.Terrain()
}

type HumanPlayer struct {
	BasePlayer
}

func NewHumanPlayer(g *Game) *HumanPlayer {
	bp := NewBasePlayer(g)
	return &HumanPlayer{*bp}
}

func (hp *HumanPlayer) Update() {
}

type AIPlayer struct {
	BasePlayer
}

func NewAIPlayer(g *Game) *AIPlayer {
	bp := NewBasePlayer(g)
	return &AIPlayer{*bp}
}

func (a *AIPlayer) Update() {
}
