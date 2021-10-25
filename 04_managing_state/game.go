package main

import (
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/spf13/viper"
)

type Game struct {
	state      GameState
	background *ebiten.Image
}

func NewGame() *Game {
	var g Game

	g.state = StateIntro

	img, _, err := ebitenutil.NewImageFromFile("assets/background.png")
	if err != nil {
		log.Fatal(err)
	}
	g.background = img

	return &g
}

func (g *Game) Update() error {
	switch g.state {
	case StateIntro:
		time.Sleep(time.Millisecond * 500)
		g.state = StateTitle

	case StateTitle:
		time.Sleep(time.Millisecond * 500)
		g.state = StateGameStart

	case StateGameStart:
		time.Sleep(time.Millisecond * 500)
		g.state = StateDealHand

	case StateDealHand:
		time.Sleep(time.Millisecond * 500)
		g.state = StateTurnStart

	case StateTurnStart:
		time.Sleep(time.Millisecond * 500)
		g.state = StateDrawPhase

	case StateDrawPhase:
		time.Sleep(time.Millisecond * 500)
		g.state = StateDiscardPhase

	case StateDiscardPhase:
		time.Sleep(time.Millisecond * 500)
		g.state = StateTurnEnd

	case StateTurnEnd:
		time.Sleep(time.Millisecond * 500)
		g.state = StateGameOver

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case StateIntro:
		PrintCentered(screen, "intro")

	case StateTitle:
		PrintCentered(screen, "title screen")

	case StateGameStart:
		PrintCentered(screen, "game start")

	case StateTurnStart, StateDealHand, StateDrawPhase, StateDiscardPhase, StateTurnEnd:
		op := ResizeTo(g.background, nil, viper.GetInt("screen_width"), viper.GetInt("screen_height"))
		screen.DrawImage(g.background, op)

	case StateGameOver:
		PrintCentered(screen, "game over")
	}
}

func (g *Game) Layout(width, height int) (int, int) {
	return viper.GetInt("screen_width"), viper.GetInt("screen_height")
}

func PrintCentered(screen *ebiten.Image, msg string) {
	bounds := text.BoundString(ttfRobotoLarge, msg)
	text.Draw(screen, msg, ttfRobotoLarge, viper.GetInt("screen_width")/2-bounds.Dx()/2, viper.GetInt("screen_height")/2-bounds.Dy()/2, color.White)
}

func ResizeTo(image *ebiten.Image, op *ebiten.DrawImageOptions, width, height int) *ebiten.DrawImageOptions {
	if op == nil {
		op = &ebiten.DrawImageOptions{}
	}

	w, h := image.Size()

	scaleW := float64(width) / float64(w)
	scaleH := float64(height) / float64(h)

	op.GeoM.Scale(scaleW, scaleH)
	return op
}
