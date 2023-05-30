package blackjack

import (
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
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("You entered:", line)
}
