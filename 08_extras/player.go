package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player interface {
	Update()
	DrawField(target *ebiten.Image)
	Hand() *Hand
	Distance() int
	Status() PlayerStatus
	Resistance() []PlayerStatus
	Terrain() string
	Receive(card Card) error
	Input() InputHandler
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

func (ps PlayerStatus) String() string {
	switch ps {
	case StatusOriented:
		return "oriented"
	case StatusWorking:
		return "working"
	case StatusEscaping:
		return "escaping"
	case StatusRecovering:
		return "recovering"
	case StatusLost:
		return "lost"
	case StatusOutOfMoney:
		return "outofmoney"
	case StatusCaptive:
		return "captive"
	case StatusSick:
		return "sick"
	default:
		return "unknown"
	}
}

type BasePlayer struct {
	g     *Game
	h     *Hand
	bs    *BattleStack
	ts    *TerrainStack
	tf    *TravelField
	df    *DefenseField
	input InputHandler
}

func NewBasePlayer(g *Game, input InputHandler, player int) *BasePlayer {
	return &BasePlayer{
		g:     g,
		h:     NewHand(),
		bs:    NewBattleStack(player),
		ts:    NewTerrainStack(player),
		tf:    NewTravelField(player),
		df:    NewDefenseField(player),
		input: input,
	}
}

func (bp *BasePlayer) Hand() *Hand {
	return bp.h
}

func (bp *BasePlayer) Distance() int {
	return bp.tf.Distance()
}

func (bp *BasePlayer) Status() PlayerStatus {
	status := bp.bs.Status()
	switch status {
	case StatusUnknown:
		status = StatusLost
		fallthrough
	case StatusLost, StatusEscaping, StatusRecovering, StatusWorking:
		hasLostResistance := false
		for _, res := range bp.Resistance() {
			if res == StatusLost {
				hasLostResistance = true
				break
			}
		}

		if hasLostResistance {
			status = StatusOriented
		}
	}

	return status
}

func (bp *BasePlayer) Resistance() []PlayerStatus {
	return bp.df.Resistance()
}

func (bp *BasePlayer) Terrain() string {
	return bp.ts.Terrain()
}

func (bp *BasePlayer) Receive(card Card) error {
	// validate card rules

	// status
	if len(card.Rules.Status) > 0 {
		validStatus := false
		for _, s := range card.Rules.Status {
			if s == bp.Status() {
				validStatus = true
				break
			}
		}

		if !validStatus {
			return fmt.Errorf("base player: can't play card %s on status %s", card.Name, bp.Status())
		}
	}

	// terrain
	if len(card.Rules.Terrains) > 0 && bp.Terrain() != "any" {
		validTerrain := false
		for _, t := range card.Rules.Terrains {
			if t == bp.Terrain() {
				validTerrain = true
				break
			}
		}

		if !validTerrain {
			return fmt.Errorf("base player: can't play card %s on terrain %s", card.Name, bp.Terrain())
		}
	}

	// resistances
	for _, res := range bp.Resistance() {
		if card.Effects.Status == res || (res == StatusLost && card.Type == "terrain") {
			return fmt.Errorf("base player: can't play card %s, player has %s resistance", card.Name, res)
		}
	}

	// travel
	if bp.Distance()+card.Effects.Distance > viper.GetInt("rules.max_distance") {
		return fmt.Errorf("base player: can't play card %s, distance overflow %d", card.Name, bp.Distance()+card.Effects.Distance)
	}

	switch card.Type {
	case "defense":
		bp.df.Put(card)
	case "battle":
		bp.bs.Put(card)
	case "travel":
		bp.tf.Put(card)
	case "terrain":
		bp.ts.Put(card)
	default:
		log.Fatalf("missing card type: %s", card.Name)
	}

	// new resistance
	if card.Effects.Resistance == bp.Status() {
		oldCard, err := bp.bs.Pop()
		if err != nil {
			// should never happen!
			log.Fatal(err)
		}

		bp.g.graveyard.Put(oldCard)
	}

	// apply AR terrain effect
	if card.Effects.Resistance == StatusLost {
		for {
			oldCard, err := bp.ts.Pop()
			if err != nil {
				log.Println(err)
				break
			}
			bp.g.graveyard.Put(oldCard)
		}
	}

	return nil
}

func (bp *BasePlayer) DrawField(target *ebiten.Image) {
	bp.ts.DrawImage(target)
	bp.bs.DrawImage(target)
	bp.tf.DrawImage(target)
	bp.df.DrawImage(target)
}

func (bp *BasePlayer) Input() InputHandler {
	return bp.input
}

type HumanPlayer struct {
	BasePlayer
}

func NewHumanPlayer(g *Game, input InputHandler, player int) *HumanPlayer {
	bp := NewBasePlayer(g, input, player)
	return &HumanPlayer{*bp}
}

func (hp *HumanPlayer) Update() {
	// nop for humans
}

type AIPlayer struct {
	BasePlayer
}

func NewAIPlayer(g *Game, player int) *AIPlayer {
	bp := NewBasePlayer(g, NewDumbHandler(), player)
	return &AIPlayer{*bp}
}

func (a *AIPlayer) Update() {
	// put some nice AI here
}
