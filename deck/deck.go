package deck

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	deck []Card
}

func (d *Deck) GetCard() (Card, error) {
	if len(d.deck) == 0 {
		return Card{}, fmt.Errorf("DECK DOESN'T HAVE CARD YET")
	}

	card := d.deck[len(d.deck)-1]
	d.deck = d.deck[:len(d.deck)-1]
	return card, nil
}

func (d *Deck) CreateDeck() (bool, error) {
	names := []string{
		"Aces", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Valet",
		"Queen", "King",
	}

	values := [][]int{
		{1, 11}, {2}, {3}, {4}, {5}, {6},
		{7}, {8}, {9}, {10}, {10}, {10},
		{10},
	}

	tickers := []string{
		"A", "2", "3", "4", "5", "6", "7",
		"8", "9", "10", "J", "Q", "K",
	}

	for i := 0; i < 13; i++ {
		d.deck = append(d.deck, Card{Name: names[i], Height: values[i], Ticker: tickers[i]})
		d.deck = append(d.deck, Card{Name: names[i], Height: values[i], Ticker: tickers[i]})
		d.deck = append(d.deck, Card{Name: names[i], Height: values[i], Ticker: tickers[i]})
		d.deck = append(d.deck, Card{Name: names[i], Height: values[i], Ticker: tickers[i]})
	}

	return true, nil
}

func (d *Deck) ShuffleDeck() (bool, error) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := len(d.deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.deck[i], d.deck[j] = d.deck[j], d.deck[i]
	}

	return true, nil
}
