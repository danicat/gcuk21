package main

import "testing"

func TestShuffle(t *testing.T) {
	d1 := NewDeck()
	d2 := NewDeck()
	d1.Shuffle(123)

	equal := true
	for i := 0; i < len(d1.cards); i++ {
		if d1.cards[i].Name != d2.cards[i].Name {
			equal = false
			break
		}
	}

	if equal {
		t.Fatal("expected decks to be different")
	}
}
