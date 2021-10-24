package main

import (
	"fmt"
	"testing"
)

func TestStateTransitions(t *testing.T) {
	tbl := []struct {
		name        string
		beforeState GameState
		inputs      []Input
		afterState  GameState
	}{
		{
			"transition intro to title",
			StateIntro,
			nil,
			StateTitle,
		},
		{
			"transition title to game start",
			StateTitle,
			nil,
			StateGameStart,
		},
		{
			"transition game start to deal hand",
			StateGameStart,
			nil,
			StateDealHand,
		},
		{
			"transition deal hand to turn start",
			StateDealHand,
			nil,
			StateTurnStart,
		},
		{
			"transition turn start to draw phase",
			StateTurnStart,
			nil,
			StateDrawPhase,
		},
		{
			"transition draw phase to game discard phase",
			StateDrawPhase,
			[]Input{KeyDefaultOrGraveyard},
			StateDiscardPhase,
		},
		{
			"transition discard phase to turn end",
			StateDiscardPhase,
			[]Input{KeyDefaultOrGraveyard},
			StateTurnEnd,
		},
		// {
		// 	"transition turn end to game over",
		// 	StateTurnEnd,
		// 	nil,
		// 	StateGameOver,
		// },
	}

	for tn, tc := range tbl {
		t.Run(fmt.Sprintf("test %d: %s", tn, tc.name), func(t *testing.T) {
			input := NewMockHandler()
			input.AppendKeys(tc.inputs)
			g := NewGame(input, 0, 0, 6)
			g.deck = NewDeck()
			g.state = tc.beforeState
			g.Update()

			if g.state != tc.afterState {
				t.Fatalf("expected %s got %s", tc.afterState, g.state)
			}
		})
	}
}

func TestDealHand(t *testing.T) {
	handSize := 6
	input := NewMockHandler()
	g := NewGame(input, 0, 0, handSize)
	g.state = StateGameStart
	g.Update()
	g.Update()

	for _, p := range g.players {
		if g.state != StateTurnStart {
			t.Fatalf("expected %s, got %s", StateTurnStart, g.state)
		}

		if len(p.Hand().cards) != handSize {
			t.Fatalf("expected %d cards, got %d", handSize, len(p.Hand().cards))
		}
	}
}

func TestSwitchPlayer(t *testing.T) {
	tbl := []struct {
		beforePlayer int
		afterPlayer  int
	}{
		{
			0,
			1,
		},
		{
			1,
			0,
		},
	}

	for _, tc := range tbl {
		g := NewGame(nil, 0, 0, 0)
		g.currentPlayer = tc.beforePlayer
		g.state = StateTurnEnd
		g.Update()
		if g.currentPlayer != tc.afterPlayer {
			t.Fatalf("expected %d, got %d", tc.afterPlayer, g.currentPlayer)
		}
	}
}
