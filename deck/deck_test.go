package deck

import (
	"reflect"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	deck := Deck{deck: make([]Card, 0)}
	deck.CreateDeck()

	if len(deck.deck) != 52 {
		t.Errorf("CreateDeck FAILED. Expect %v, got %v", 52, len(deck.deck))
	}

	lastCard := Card{Name: "King", Height: []int{10}, Ticker: "K"}

	if deck.deck[51].Name != lastCard.Name {
		t.Errorf("CreateDeck FAILED. Expect %v, got %v", "King", deck.deck[51].Name)
	}
}

func TestGetCard(t *testing.T) {
	deck := Deck{deck: make([]Card, 0)}
	_, err := deck.GetCard()

	if err == nil {
		t.Errorf("GetCard FAILED. Expect %v, got %v", "DECK DOESN'T HAVE CARD YET", err)
	}

	deck.CreateDeck()
	card, _ := deck.GetCard()

	if card.Name != "King" {
		t.Errorf("GetCard FAILED. Expect %v, got %v", "King", card.Name)
	}
}

func TestShuffleDeck(t *testing.T) {
	firstDeck := Deck{deck: make([]Card, 0)}
	firstDeck.CreateDeck()

	if len(firstDeck.deck) != 52 {
		t.Errorf("ShuffleDeck FAILED. Expect %v, got %v", 52, len(firstDeck.deck))
	}

	lastCard := Card{Name: "King", Height: []int{10}, Ticker: "K"}

	if firstDeck.deck[51].Name != lastCard.Name {
		t.Errorf("ShuffleDeck FAILED. Expect %v, got %v", "King", firstDeck.deck[51].Name)
	}

	secondDeck := Deck{deck: make([]Card, 0)}
	secondDeck.CreateDeck()
	secondDeck.ShuffleDeck()

	if reflect.DeepEqual(firstDeck.deck, secondDeck.deck) {
		t.Errorf("ShuffleDeck FAILED. Expect %v, got %v", false, true)
	}
}
