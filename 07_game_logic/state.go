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
		return "Unknown"
	case StateIntro:
		return "Intro"
	case StateTitle:
		return "Title"
	case StateGameStart:
		return "Game Start"
	case StateDealHand:
		return "Deal"
	case StateTurnStart:
		return "Turn Start"
	case StateDrawPhase:
		return "Draw"
	case StateDiscardPhase:
		return "Discard"
	case StateTurnEnd:
		return "Turn End"
	case StateGameOver:
		return "Game Over"
	default:
		return ""
	}
}
