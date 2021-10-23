package main

type GameState int

const (
	StateUnknown GameState = iota
	StateIntro
	StateTitle
	StateGameStart
	StateDealHand
	StateTurnStart
	StateDrawPhase
	StateDiscardPhase
	StateTurnEnd
	StateGameOver
)

func (s GameState) String() string {
	switch s {
	case StateUnknown:
		return "StateUnknown"
	case StateIntro:
		return "StateIntro"
	case StateTitle:
		return "StateTitle"
	case StateGameStart:
		return "StateGameStart"
	case StateDealHand:
		return "StateDealHand"
	case StateTurnStart:
		return "StateTurnStart"
	case StateDrawPhase:
		return "StateDrawPhase"
	case StateDiscardPhase:
		return "StateDiscardPhase"
	case StateTurnEnd:
		return "StateTurnEnd"
	case StateGameOver:
		return "StateGameOver"
	default:
		return ""
	}
}
