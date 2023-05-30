package blackjack

import (
	"blackjack-go/player"
	"bufio"
	"fmt"
	"os"
)

type Blackjack struct{}

func (b Blackjack) Run() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\x1B[2J\x1B[1;1H")
	fmt.Print("Welcome to BlackJack -_-")

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

	human := player.Human{Player: {name}}
	// let mut human: Human = Human::new(name, vec![], balance, 0, 0)

	fmt.Print("Could you choose a name for a bot? ")
	computer, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// let mut computer = Bot::new(bot, vec![], balance, 0, 0)
	// computer.welcome_message()

	fmt.Println("So let's play? ")
	fmt.Print("\x1B[2J\x1B[1;1H")
}
