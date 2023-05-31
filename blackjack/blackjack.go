package blackjack

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"blackjack-go/deck"
	"blackjack-go/player"
)

type Blackjack struct{}

func (b Blackjack) Run() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\x1B[2J\x1B[1;1H")
	fmt.Println("Welcome to BlackJack -_-")

	fmt.Print("First, What is your name? ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Print("How many money have you today? ")
	balance, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	balance_in_float, err := strconv.ParseFloat(strings.TrimSpace(balance), 64)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	human := player.Human{Player: player.Player{}}
	human.Player.New(name, make([]string, 0), balance_in_float, 0, 0)

	fmt.Print("Could you choose a name for a bot? ")
	computerName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	computer := player.Bot{Player: player.Player{}}
	computer.Player.New(computerName, make([]string, 0), balance_in_float, 0, 0)
	computer.WelcomeMessage()

	fmt.Println("So let's play? ")
	fmt.Print("\x1B[2J\x1B[1;1H")

	deck := deck.Deck{}
	game := false

	for i := 0; ; i++ {
		if !game {
			human.ClearHand()
			computer.ClearHand()
			deck.CreateDeck()
			deck.ShuffleDeck()
			fmt.Print("\x1B[2J\x1B[1;1H")

			check_balance_player, _ := b.checkBalance(human.Player)
			check_balance_computer, _ := b.checkBalance(computer.Player)

			if check_balance_player || check_balance_computer {
				game = true
				continue
			}

			fmt.Print("\x1B[2J\x1B[1;1H")
		}
	}
}

func (b Blackjack) checkBalance(player player.Player) (bool, error) {
	if balance, err := player.GetBalance(); err != nil && balance == 0.0 {
		name, _ := player.GetName()
		fmt.Println("{}, you haven't balance", name)
		return true, nil
	} else {
		return false, nil
	}
}
