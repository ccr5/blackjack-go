package player

import "fmt"

type Bot struct {
	Player
}

func (b Bot) WelcomeMessage() {
	name, _ := b.GetName()
	fmt.Println("Hi o/ \nMy name is {} and I'll play with you", name)
	fmt.Println("Good luck!")
	fmt.Println("Aaah, I have the same balance for a fair play :)")
}

func (b Bot) PlayGame() (bool, error) {

	if hand, err := b.GetHand(); err != nil && len(hand) == 0 {
		return true, nil
	} else {
		// count_a := 0
		// sum_hand := 0
		hand, err := b.GetHand()
		if err != nil {
			return false, err
		}

		for i := 0; i < len(hand); i++ {
			fmt.Print(i)
		}

		return true, nil
	}
}
