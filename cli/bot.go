package main

import "fmt"

type Bot struct {
	Player
}

func (b Bot) WelcomeMessage() {
	fmt.Printf("Hi o/ \nMy name is %v and I'll play with you\n", b.Name)
	fmt.Println("Good luck!")
	fmt.Println("Aaah, I have the same balance for a fair play :)")
}

func (b Bot) PlayGame() (bool, error) {

	if len(b.Hand) == 0 {
		return true, nil
	} else {
		count_a := 0
		sum_hand := 0

		for _, card := range b.Hand {
			if card.Ticker == "A\n" {
				count_a += 1

				if count_a >= 1 && sum_hand > 10 {
					sum_hand += 1 * count_a
				} else {
					sum_hand += 11
				}

			} else {
				sum_hand += card.Height[0]
			}
		}

		return sum_hand < 21, nil
	}
}
