package main

import (
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/spf13/viper"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
	state     GameState
	deck      *Deck
	graveyard *Stack
	players   []Player
	current   int
	input     InputHandler

	// Configuration
	screenWidth, screenHeight int
	handSize                  int

	// Graphics
	background *ebiten.Image
	op         *ebiten.DrawImageOptions
}

func NewGame(input InputHandler, screenWidth, screenHeight, handSize int) *Game {
	var g Game

	g.screenWidth = screenWidth
	g.screenHeight = screenHeight
	g.handSize = handSize
	g.state = StateIntro

	hp := NewHumanPlayer(&g)
	ap := NewAIPlayer(&g)

	g.players = append(g.players, hp, ap)

	img, _, err := ebitenutil.NewImageFromFile(viper.GetString("field.background"))
	if err != nil {
		log.Fatal(err)
	}
	g.background = img
	g.op = ResizeTo(img, nil, screenWidth, screenHeight)

	g.input = input

	return &g
}

func (g *Game) NextPlayer() int {
	return (g.current + 1) % 2
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
		g.deck = NewDeck(0, 0, 0, 0)
		g.deck.Shuffle(time.Now().UnixNano())
		g.graveyard = &Stack{}
		g.current = 0
		g.state = StateDealHand

	case StateDealHand:
		for i := 0; i < g.handSize; i++ {
			for _, p := range g.players {
				c, err := g.deck.DrawCard()
				if err != nil {
					log.Fatal(err)
				}
				p.Hand().Put(c)
			}
		}
		g.state = StateTurnStart

	case StateTurnStart:
		time.Sleep(time.Millisecond * 200)
		g.state = StateDrawPhase

	case StateDrawPhase:
		switch key := g.input.Read(); key {
		case KeyDefaultOrGraveyard:
			g.state = StateDiscardPhase
		}

	case StateDiscardPhase:
		switch key := g.input.Read(); key {
		case KeyDefaultOrGraveyard:
			g.state = StateTurnEnd
		}

	case StateTurnEnd:
		time.Sleep(time.Millisecond * 1000)
		g.state = StateGameOver

	case StateGameOver:

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case StateIntro:
		g.PrintCentered(screen, "intro state")

	case StateTitle:
		g.PrintCentered(screen, "title screen")

	case StateGameStart:
		g.PrintCentered(screen, "game start")

	case StateTurnStart, StateDealHand, StateDrawPhase, StateDiscardPhase, StateTurnEnd:
		screen.DrawImage(g.background, g.op)
		g.deck.DrawImage(screen)

	case StateGameOver:
		g.PrintCentered(screen, "game over")
	}
}

func (g *Game) Layout(width, height int) (int, int) {
	return g.screenWidth, g.screenHeight
}

func (g *Game) PrintCentered(screen *ebiten.Image, msg string) {
	bounds := text.BoundString(ttfRoboto, msg)
	text.Draw(screen, msg, ttfRoboto, g.screenWidth/2-bounds.Dx()/2, g.screenHeight/2-bounds.Dy()/2, color.White)
}
