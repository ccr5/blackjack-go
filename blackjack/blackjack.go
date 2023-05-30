package blackjack

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

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

	fmt.Println(balance)

	balance_in_float, err := strconv.ParseFloat(balance, 32)
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
}
