package main

import (
	"fmt"
	"log"
	"testing"
)

func TestReceive(t *testing.T) {
	tbl := []struct {
		name          string
		beforeCards   []string
		afterStatus   PlayerStatus
		afterDistance int
		afterTerrain  string
	}{
		// No cards
		{
			"should be lost",
			[]string{},
			StatusLost,
			0,
			"any",
		},
		// Green card effects
		{
			"should be oriented",
			[]string{"orientation"},
			StatusOriented,
			0,
			"any",
		},
		{
			"should be oriented",
			[]string{"orientation", "lost", "orientation"},
			StatusOriented,
			0,
			"any",
		},
		{
			"should be escaping",
			[]string{"orientation", "hostiles", "escape"},
			StatusEscaping,
			0,
			"any",
		},
		{
			"should be working",
			[]string{"orientation", "outofmoney", "work"},
			StatusWorking,
			0,
			"any",
		},
		{
			"should be recovering",
			[]string{"orientation", "epidemic", "remedy"},
			StatusRecovering,
			0,
			"any",
		},
		// Red card effects
		{
			"should be lost",
			[]string{"orientation", "lost"},
			StatusLost,
			0,
			"any",
		},
		{
			"should be captive",
			[]string{"orientation", "hostiles"},
			StatusCaptive,
			0,
			"any",
		},
		{
			"should be out of money",
			[]string{"orientation", "outofmoney"},
			StatusOutOfMoney,
			0,
			"any",
		},
		{
			"should be sick",
			[]string{"orientation", "epidemic"},
			StatusSick,
			0,
			"any",
		},
		// Should be oriented (RA)
		{
			"should be oriented (RA)",
			[]string{"routes", "hostiles", "escape"},
			StatusOriented,
			0,
			"any",
		},
		{
			"should be oriented (RA)",
			[]string{"routes", "outofmoney", "work"},
			StatusOriented,
			0,
			"any",
		},
		{
			"should be oriented (RA)",
			[]string{"routes", "epidemic", "remedy"},
			StatusOriented,
			0,
			"any",
		},
		// Blue card status cancel
		{
			"should be oriented (RA)",
			[]string{"routes"},
			StatusOriented,
			0,
			"any",
		},
		{
			"should be oriented (RA)",
			[]string{"orientation", "lost", "routes"},
			StatusOriented,
			0,
			"any",
		},
		// Blue card immunities
		{
			"should be immune to lost",
			[]string{"orientation", "routes", "lost"},
			StatusOriented,
			0,
			"any",
		},
		{
			"should be immune to captive",
			[]string{"orientation", "diplomacy", "hostiles"},
			StatusOriented,
			0,
			"any",
		},
		{
			"should be immune to out of money",
			[]string{"orientation", "riches", "outofmoney"},
			StatusOriented,
			0,
			"any",
		},
		{
			"should be immune to sick",
			[]string{"orientation", "health", "epidemic"},
			StatusOriented,
			0,
			"any",
		},
		// Yellow cards
		{
			"should be desert",
			[]string{"desert"},
			StatusLost,
			0,
			"desert",
		},
		{
			"should be desert",
			[]string{"civilization", "desert"},
			StatusLost,
			0,
			"desert",
		},
		{
			"should be civilization",
			[]string{"desert", "civilization"},
			StatusLost,
			0,
			"civilization",
		},
		{
			"should be savage land",
			[]string{"desert", "civilization", "savage_land"},
			StatusLost,
			0,
			"savage_land",
		},
		{
			"should be sea",
			[]string{"desert", "civilization", "savage_land", "sea"},
			StatusLost,
			0,
			"sea",
		},
		{
			"should be no terrain (RA)",
			[]string{"desert", "civilization", "sea", "savage_land", "routes"},
			StatusOriented,
			0,
			"any",
		},
		{
			"should be no terrain (RA)",
			[]string{"routes", "desert", "savage_land", "sea", "civilization"},
			StatusOriented,
			0,
			"any",
		},
		// White cards
		{
			"1000",
			[]string{"orientation", "1000"},
			StatusOriented,
			1000,
			"any",
		},
		{
			"1000",
			[]string{"orientation", "desert", "1000"},
			StatusOriented,
			1000,
			"desert",
		},
		{
			"1000",
			[]string{"orientation", "savage_land", "1000"},
			StatusOriented,
			1000,
			"savage_land",
		},
		{
			"1000",
			[]string{"orientation", "civilization", "1000"},
			StatusOriented,
			1000,
			"civilization",
		},
		{
			"0",
			[]string{"orientation", "sea", "1000"},
			StatusOriented,
			0,
			"sea",
		},
		{
			"0",
			[]string{"orientation", "desert", "lost", "1000"},
			StatusLost,
			0,
			"desert",
		},
		{
			"1000",
			[]string{"orientation", "sea", "lost", "routes", "1000"},
			StatusOriented,
			1000,
			"any",
		},
	}

	for _, tc := range tbl {
		t.Run(fmt.Sprintf(tc.name), func(t *testing.T) {
			g := NewGame(nil, 0, 0, 6)
			p := g.players[0]
			g.state = StateGameStart
			g.Update()

			for _, c := range tc.beforeCards {
				card, ok := Cards[c]
				if !ok {
					t.Fatalf("card %s not found", c)
				}

				err := p.Receive(card)
				if err != nil {
					log.Println(err)
				}
			}

			if p.Status() != tc.afterStatus {
				t.Errorf("status should be %s, got %s", tc.afterStatus, p.Status())
			}

			if p.Distance() != tc.afterDistance {
				t.Errorf("distance should be %d, got %d", tc.afterDistance, p.Distance())
			}

			if p.Terrain() != tc.afterTerrain {
				t.Errorf("terrain should be %s, got %s", tc.afterTerrain, p.Terrain())
			}
		})
	}

}
