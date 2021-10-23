package main

import (
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/spf13/viper"
)

type Game struct {
	state GameState
}

func NewGame() *Game {
	var g Game

	g.state = StateIntro

	return &g
}

func (g *Game) Update() error {
	switch g.state {
	case StateIntro:
		time.Sleep(time.Millisecond * 200)
		g.state = StateTitle

	case StateTitle:
		time.Sleep(time.Millisecond * 200)
		g.state = StateGameStart

	case StateGameStart:
		time.Sleep(time.Millisecond * 200)
		g.state = StateDealHand

	case StateDealHand:
		time.Sleep(time.Millisecond * 200)
		g.state = StateTurnStart

	case StateTurnStart:
		time.Sleep(time.Millisecond * 200)
		g.state = StateDrawPhase

	case StateDrawPhase:
		time.Sleep(time.Millisecond * 200)
		g.state = StateDiscardPhase

	case StateDiscardPhase:
		time.Sleep(time.Millisecond * 200)
		g.state = StateTurnEnd

	case StateTurnEnd:
		time.Sleep(time.Millisecond * 200)
		g.state = StateGameOver

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case StateIntro:
		PrintCentered(screen, "intro state")

	case StateTitle:
		PrintCentered(screen, "title screen")

	case StateGameStart:
		PrintCentered(screen, "game start")

	case StateTurnStart, StateDealHand, StateDrawPhase, StateDiscardPhase, StateTurnEnd:
		screen.Fill(Colors["khaki"])

	case StateGameOver:
		PrintCentered(screen, "game over")
	}
}

func (g *Game) Layout(width, height int) (int, int) {
	return viper.GetInt("screen_width"), viper.GetInt("screen_height")
}

func PrintCentered(screen *ebiten.Image, msg string) {
	bounds := text.BoundString(ttfRoboto, msg)
	text.Draw(screen, msg, ttfRoboto, viper.GetInt("screen_width")/2-bounds.Dx()/2, viper.GetInt("screen_height")/2-bounds.Dy()/2, color.White)
}
