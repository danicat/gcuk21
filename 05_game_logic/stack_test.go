package main

import "testing"

func TestPut(t *testing.T) {
	s := Stack{}
	card := Card{Name: "foo"}
	s.Put(card)
	result, err := s.Pop()
	if err != nil {
		t.Fatal(err)
	}
	if result.Name != card.Name {
		t.Fatalf("expected %s, got %s", card.Name, result.Name)
	}
}

func TestBattleStackPut(t *testing.T) {
	s := NewBattleStack(0)
	card := Card{Name: "foo"}
	s.Put(card)
	result, err := s.Pop()
	if err != nil {
		t.Fatal(err)
	}
	if result.Name != card.Name {
		t.Fatalf("expected %s, got %s", card.Name, result.Name)
	}
}

func TestTerrainStackPut(t *testing.T) {
	s := NewTerrainStack(0)
	card := Card{Name: "foo"}
	s.Put(card)
	result, err := s.Pop()
	if err != nil {
		t.Fatal(err)
	}
	if result.Name != card.Name {
		t.Fatalf("expected %s, got %s", card.Name, result.Name)
	}
}
