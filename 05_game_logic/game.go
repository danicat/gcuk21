package main

import (
	"fmt"
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
	state         GameState
	deck          *Deck
	graveyard     *Graveyard
	players       []Player
	currentPlayer int
	emptyHand     bool

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

	hp := NewHumanPlayer(&g, input, 1)
	ap := NewAIPlayer(&g, 2)
	g.players = append(g.players, hp, ap)

	img, _, err := ebitenutil.NewImageFromFile(viper.GetString("field.background"))
	if err != nil {
		log.Fatal(err)
	}
	g.background = img
	g.op = ResizeTo(img, nil, screenWidth, screenHeight)

	return &g
}

func (g *Game) NextPlayer() int {
	return (g.currentPlayer + 1) % 2
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
		g.deck = NewDeck()
		g.deck.Shuffle(time.Now().UnixNano())
		g.graveyard = NewGraveyard()
		g.currentPlayer = 0
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
		switch key := g.players[g.currentPlayer].Input().Read(); key {
		case KeyDefaultOrGraveyard:
			card, err := g.deck.DrawCard()
			if err != nil {
				// deck is empty, but we allow the game to continue until we have no cards left on hand
				log.Println(err)
				g.state = StateDiscardPhase
				return nil
			}
			g.players[g.currentPlayer].Hand().Put(card)
			g.state = StateDiscardPhase
		}

	case StateDiscardPhase:
		switch key := g.players[g.currentPlayer].Input().Read(); key {
		case KeyDefaultOrGraveyard:
			card, err := g.players[g.currentPlayer].Hand().Discard()
			if err != nil {
				log.Println(err)
				// no cards left in hand, is it over?
				if g.emptyHand {
					// both players have no cards left, wrap it up
					g.state = StateGameOver
				}
				g.emptyHand = true
				g.state = StateTurnEnd
				return nil
			}

			if card.Rules.Target == "opponent" {
				g.Play(g.players[g.currentPlayer], g.players[g.NextPlayer()], card)
			} else {
				// self or any
				g.Play(g.players[g.currentPlayer], g.players[g.currentPlayer], card)
			}

			g.state = StateTurnEnd

		case KeyOpponentOrGraveyard:
			card, err := g.players[g.currentPlayer].Hand().Discard()
			if err != nil {
				log.Println(err)
				// no cards left in hand, is it over?
				if g.emptyHand {
					// both players have no cards left, wrap it up
					g.state = StateGameOver
				}
				g.emptyHand = true
				g.state = StateTurnEnd
				return nil
			}

			if card.Rules.Target == "self" {
				g.graveyard.Put(card)
			} else {
				g.Play(g.players[g.currentPlayer], g.players[g.NextPlayer()], card)
			}

			g.state = StateTurnEnd

		case KeyGraveyard:
			card, err := g.players[g.currentPlayer].Hand().Discard()
			if err != nil {
				log.Println(err)
				// no cards left in hand, is it over?
				if g.emptyHand {
					// both players have no cards left, wrap it up
					g.state = StateGameOver
				}
				g.emptyHand = true
				g.state = StateTurnEnd
				return nil
			}

			g.graveyard.Put(card)

			g.state = StateTurnEnd

		case KeyLeft:
			if viper.GetBool("reverse_left_right") {
				g.players[g.currentPlayer].Hand().Right()
			} else {
				g.players[g.currentPlayer].Hand().Left()
			}

		case KeyRight:
			if viper.GetBool("reverse_left_right") {
				g.players[g.currentPlayer].Hand().Left()
			} else {
				g.players[g.currentPlayer].Hand().Right()
			}
		}

	case StateTurnEnd:
		if g.players[g.currentPlayer].Distance() == viper.GetInt("rules.max_distance") {
			log.Printf("player %d cleared the game!", g.currentPlayer)
			time.Sleep(time.Second)
			g.state = StateGameOver
			return nil
		}
		g.currentPlayer = g.NextPlayer()
		g.state = StateTurnStart

	case StateGameOver:

	}

	return nil
}

func (g *Game) Play(from, to Player, card Card) {
	if to == nil {
		g.graveyard.Put(card)
		return
	}

	err := to.Receive(card)
	if err != nil {
		log.Println(err)
		g.graveyard.Put(card)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case StateIntro:
		g.PrintCentered(screen, "intro state", color.White)

	case StateTitle:
		g.PrintCentered(screen, "title screen", color.White)

	case StateGameStart:
		g.PrintCentered(screen, "game start", color.White)

	case StateTurnStart, StateDealHand, StateDrawPhase, StateDiscardPhase, StateTurnEnd:
		screen.DrawImage(g.background, g.op)

		score1 := fmt.Sprintf("Travel distance: %d Km", g.players[0].Distance())
		score2 := fmt.Sprintf("Travel distance: %d Km", g.players[1].Distance())
		text.Draw(screen, score1, ttfRoboto, viper.GetInt("layout.player1.score.startX"), viper.GetInt("layout.player1.score.startY"), color.Black)
		text.Draw(screen, score2, ttfRoboto, viper.GetInt("layout.player2.score.startX"), viper.GetInt("layout.player2.score.startY"), color.Black)

		g.deck.DrawImage(screen)
		g.graveyard.DrawImage(screen)

		for _, p := range g.players {
			p.DrawField(screen)
		}

		if g.currentPlayer == 0 || !viper.GetBool("hide_cpu_hand") {
			g.players[g.currentPlayer].Hand().DrawImage(screen)
		}

		text.Draw(screen, fmt.Sprintf("Current Phase: Player %d - %s", g.currentPlayer+1, g.state), ttfRoboto, viper.GetInt("layout.info.startX"), viper.GetInt("layout.info.startY"), color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF})

	case StateGameOver:
		g.PrintCentered(screen, "GAME OVER", color.White)
	}
}

func (g *Game) Layout(width, height int) (int, int) {
	return g.screenWidth, g.screenHeight
}

func (g *Game) PrintCentered(screen *ebiten.Image, msg string, clr color.Color) {
	bounds := text.BoundString(ttfRobotoLarge, msg)
	text.Draw(screen, msg, ttfRobotoLarge, g.screenWidth/2-bounds.Dx()/2, g.screenHeight/2-bounds.Dy()/2, clr)
}
