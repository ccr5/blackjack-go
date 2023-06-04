package player

import (
	"testing"

	"github.com/ccr5/blackjack-go/deck"
)

func TestAddCardToHand(t *testing.T) {
	player := Player{
		Name:    "name",
		Hand:    make([]deck.Card, 0),
		Balance: 0.0,
		Wins:    0,
		Defeats: 0,
	}

	if len(player.Hand) != 0 {
		t.Errorf("AddCardToHand FAILED. Expected %v, got %v", 0, len(player.Hand))
	}

	card := deck.Card{Name: "Ten", Height: []int{10}, Ticker: "10"}
	player.AddCardToHand(card)
	if len(player.Hand) != 1 {
		t.Errorf("AddCardToHand FAILED. Expected %v, got %v", 1, len(player.Hand))
	}
}

func TestClearHand(t *testing.T) {
	player := Player{
		Name:    "name",
		Hand:    make([]deck.Card, 0),
		Balance: 0.0,
		Wins:    0,
		Defeats: 0,
	}

	if len(player.Hand) != 0 {
		t.Errorf("ClearHand FAILED. Expected %v, got %v", 0, len(player.Hand))
	}

	card := deck.Card{Name: "Ten", Height: []int{10}, Ticker: "10"}
	player.AddCardToHand(card)
	if len(player.Hand) != 1 {
		t.Errorf("ClearHand FAILED. Expected %v, got %v", 1, len(player.Hand))
	}

	player.ClearHand()
	if len(player.Hand) != 0 {
		t.Errorf("ClearHand FAILED. Expected %v, got %v", 0, len(player.Hand))
	}
}

func TestDeposit(t *testing.T) {
	player := Player{
		Name:    "name",
		Hand:    make([]deck.Card, 0),
		Balance: 0.0,
		Wins:    0,
		Defeats: 0,
	}

	if player.Balance != 0.0 {
		t.Errorf("Deposit FAILED. Expected %v, got %v", 0.0, player.Balance)
	}

	player.Deposit(100.0)
	if player.Balance != 100.0 {
		t.Errorf("Deposit FAILED. Expected %v, got %v", 100.0, player.Balance)
	}
}

func TestWithdraw(t *testing.T) {
	player := Player{
		Name:    "name",
		Hand:    make([]deck.Card, 0),
		Balance: 100.0,
		Wins:    0,
		Defeats: 0,
	}

	if player.Balance != 100.0 {
		t.Errorf("Withdraw FAILED. Expected %v, got %v", 100.0, player.Balance)
	}

	player.Withdraw(10.0)
	if player.Balance != 90.0 {
		t.Errorf("Withdraw FAILED. Expected %v, got %v", 90.0, player.Balance)
	}
}

func TestAddWin(t *testing.T) {
	player := Player{
		Name:    "name",
		Hand:    make([]deck.Card, 0),
		Balance: 100.0,
		Wins:    0,
		Defeats: 0,
	}

	if player.Wins != 0 {
		t.Errorf("AddWin FAILED. Expected %v, got %v", 0, player.Wins)
	}

	player.AddWin()
	if player.Wins != 1 {
		t.Errorf("AddWin FAILED. Expected %v, got %v", 1, player.Wins)
	}
}

func TestAddDefeats(t *testing.T) {
	player := Player{
		Name:    "name",
		Hand:    make([]deck.Card, 0),
		Balance: 100.0,
		Wins:    0,
		Defeats: 0,
	}

	if player.Defeats != 0 {
		t.Errorf("AddDefeats FAILED. Expected %v, got %v", 0, player.Defeats)
	}

	player.AddDefeats()
	if player.Defeats != 1 {
		t.Errorf("AddDefeats FAILED. Expected %v, got %v", 1, player.Defeats)
	}
}

func TestShowHand(t *testing.T) {
	t.Log("ShowHand Ok")
}

func TestShowInfo(t *testing.T) {
	t.Log("ShowInfo Ok")
}
