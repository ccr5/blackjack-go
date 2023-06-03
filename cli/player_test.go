package main

import "testing"

func TestAddCardToHand(t *testing.T) {
	player := Player{
		Name:    "name",
		Hand:    make([]Card, 0),
		Balance: 0.0,
		Wins:    0,
		Defeats: 0,
	}

	if len(player.Hand) != 0 {
		t.Errorf("AddCardToHand FAILED. Expected %v, got %v", 0, len(player.Hand))
	}

	card := Card{Name: "Ten", Height: []int{10}, Ticker: "10"}

	player.AddCardToHand(card)

	if len(player.Hand) != 1 {
		t.Errorf("AddCardToHand FAILED. Expected %v, got %v", 1, len(player.Hand))
	}
}
