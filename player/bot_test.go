package player

import (
	"testing"

	"github.com/ccr5/blackjack-go/deck"
)

func TestWelcomeMessage(t *testing.T) {
	t.Log("WelcomeMessage Ok")
}

func TestPlayGame(t *testing.T) {
	bot := Bot{
		Player: Player{
			Name:    "name",
			Hand:    make([]deck.Card, 0),
			Balance: 0.0,
			Wins:    0,
			Defeats: 0,
		},
	}

	firstResult, _ := bot.PlayGame()
	if !firstResult {
		t.Errorf("PlayGame FAILED. Expect %v, got %v", true, firstResult)
	}

	card := deck.Card{Name: "Ten", Height: []int{10}, Ticker: "10"}

	bot.AddCardToHand(card)
	secondResult, _ := bot.PlayGame()
	if !secondResult {
		t.Errorf("PlayGame FAILED. Expect %v, got %v", true, secondResult)
	}

	bot.AddCardToHand(card)
	thirdResult, _ := bot.PlayGame()
	if !thirdResult {
		t.Errorf("PlayGame FAILED. Expect %v, got %v", true, thirdResult)
	}

	bot.AddCardToHand(card)
	lastResult, _ := bot.PlayGame()
	if lastResult {
		t.Errorf("PlayGame FAILED. Expect %v, got %v", true, lastResult)
	}
}
