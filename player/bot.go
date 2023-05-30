package player

import "fmt"

type Bot struct {
	Player
}

func (b Bot) WelcomeMessage() {
	fmt.Println("Hi o/ \nMy name is {} and I'll play with you", self.player.get_name())
	fmt.Println("Good luck!")
	fmt.Println("Aaah, I have the same balance for a fair play :)")
}

func (b Bot) play_game() (bool, error) {

	if hand, err := b.GetHand(); !err && len(hand) == 0 {
		return true, nil
	} else {
		let mut count_a = 0
		let mut sum_hand = 0

		for card in self.player.get_hand() {
			if card.get_ticker() == "A" {
				count_a += 1

				if count_a.ge(&1) && sum_hand.gt(&10) {
					sum_hand += 1 * count_a
				} else  {
					sum_hand += 11
				}

			} else {
				sum_hand += card.get_height()[0]
			}
		}
			

	sum_hand < 21
	}
}