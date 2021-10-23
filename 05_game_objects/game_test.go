package main

import (
	"fmt"
	"testing"
)

func TestStateTransitions(t *testing.T) {
	tbl := []struct {
		name        string
		beforeState GameState
		afterState  GameState
	}{
		{
			"transition intro to title",
			StateIntro,
			StateTitle,
		},
		{
			"transition title to game start",
			StateTitle,
			StateGameStart,
		},
		{
			"transition game start to deal hand",
			StateGameStart,
			StateDealHand,
		},
		{
			"transition deal hand to turn start",
			StateDealHand,
			StateTurnStart,
		},
		{
			"transition turn start to draw phase",
			StateTurnStart,
			StateDrawPhase,
		},
		{
			"transition draw phase to game discard phase",
			StateDrawPhase,
			StateDiscardPhase,
		},
		{
			"transition discard phase to turn end",
			StateDiscardPhase,
			StateTurnEnd,
		},
		{
			"transition turn end to game over",
			StateTurnEnd,
			StateGameOver,
		},
	}

	for tn, tc := range tbl {
		t.Run(fmt.Sprintf("test %d: %s", tn, tc.name), func(t *testing.T) {
			g := NewGame(0, 0, 0)
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
	g := NewGame(0, 0, handSize)
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
