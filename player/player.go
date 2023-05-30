package player

import "fmt"

type Player struct {
	name    string
	hand    []string
	balance float64
	wins    int
	defeats int
}

type PlayerType interface {
	New(name string, hand []string, balance float64, wins int, defeats int) (Player, error)
	GetName() (string, error)
	GetHand() ([]string, error)
	AddCardToHand(card string) (bool, error)
	ClearHand() (bool, error)
	GetBalance() (float32, error)
	Deposit(value float32) (bool, error)
	Withdraw(value float32) (bool, error)
	GetWins() (int, error)
	addWin() (bool, error)
	GetDefeats() (int, error)
	AddDefeats() (bool, error)
	ShowHand()
}

func (p Player) New(name string, hand []string, balance float64, wins int, defeats int) (Player, error) {
	newPlayer := Player{
		name:    name,
		hand:    hand,
		balance: balance,
		wins:    wins,
		defeats: defeats,
	}

	return newPlayer, nil
}

func (p Player) GetName() (string, error) {
	return p.name, nil
}

func (p Player) GetHand() ([]string, error) {
	return p.hand, nil
}

func (p *Player) AddCardToHand(card string) (bool, error) {
	p.hand = append(p.hand, card)
	return true, nil
}

func (p *Player) ClearHand() (bool, error) {
	p.hand = make([]string, 1)
	return true, nil
}

func (p Player) GetBalance() (float64, error) {
	return p.balance, nil
}

func (p *Player) Deposit(value float64) (bool, error) {
	p.balance += value
	return true, nil
}

func (p *Player) Withdraw(value float64) (bool, error) {
	p.balance -= value
	return true, nil
}

func (p Player) GetWins() (int, error) {
	return p.wins, nil
}

func (p *Player) addWin() (bool, error) {
	p.wins += 1
	return true, nil
}

func (p Player) GetDefeats() (int, error) {
	return p.defeats, nil

}

func (p *Player) AddDefeats() (bool, error) {
	p.defeats += 1
	return true, nil
}

func (p Player) ShowHand() {
	for i := 0; i < len(p.hand); i++ {
		fmt.Print(p.hand[i])
	}
}
