package player

import (
	"fmt"

	"github.com/ccr5/blackjack-go/deck"
)

type Player struct {
	Name    string
	Hand    []deck.Card
	Balance float64
	Wins    int
	Defeats int
}

type PlayerType interface {
	AddCardToHand(card string) (bool, error)
	ClearHand() (bool, error)
	Deposit(value float32) (bool, error)
	Withdraw(value float32) (bool, error)
	addWin() (bool, error)
	AddDefeats() (bool, error)
	ShowHand()
}

func (p *Player) AddCardToHand(card deck.Card) (bool, error) {
	p.Hand = append(p.Hand, card)
	return true, nil
}

func (p *Player) ClearHand() (bool, error) {
	p.Hand = make([]deck.Card, 0)
	return true, nil
}

func (p *Player) Deposit(value float64) (bool, error) {
	p.Balance += value
	return true, nil
}

func (p *Player) Withdraw(value float64) (bool, error) {
	p.Balance -= value
	return true, nil
}

func (p *Player) AddWin() (bool, error) {
	p.Wins += 1
	return true, nil
}

func (p *Player) AddDefeats() (bool, error) {
	p.Defeats += 1
	return true, nil
}

func (p Player) ShowHand() {
	for i := 0; i < len(p.Hand); i++ {
		p.Hand[i].ShowCard()
	}
}

func (p Player) ShowInfo() {
	if p.Balance > 0 {
		fmt.Printf("%v, your balance is: %v\n", p.Name, p.Balance)
	} else if p.Balance == 0 {
		fmt.Printf("%v, you haven't balance, thanks for play BlackJack\n", p.Name)
	} else {
		fmt.Printf("%v, you're owe %v\n", p.Name, p.Balance)
	}
}
