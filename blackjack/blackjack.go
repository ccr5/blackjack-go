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
	human.Player.New(name, make([]deck.Card, 0), balance_in_float, 0, 0)

	fmt.Print("Could you choose a name for a bot? ")
	computerName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	computer := player.Bot{Player: player.Player{}}
	computer.Player.New(computerName, make([]deck.Card, 0), balance_in_float, 0, 0)
	computer.WelcomeMessage()

	fmt.Println("So let's play? ")
	fmt.Print("\x1B[2J\x1B[1;1H")

	deck := deck.Deck{}
	game := false

	for {
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

			bets, err := b.bet(&human.Player, &computer.Player)
			if err != nil {
				panic("Bet error")
			}

			ask_player := false
			for {

				if ask_player {
					break
				}

				if hand, err := human.GetHand(); err == nil && len(hand) == 0 {
					fmt.Println("You don't have any card in your hand")
				} else {
					human.ShowHand()
				}

				fmt.Print("Would you like one more card? ")
				more_card, err := reader.ReadString('\n')
				if err != nil {
					panic("Error to check more card")
				}

				more_card_upper := strings.ToUpper(strings.TrimSpace(more_card))

				if more_card_upper == "Y" {
					card, _ := deck.GetCard()
					human.AddCardToHand(card)
				} else {
					ask_player = true
				}
			}

			bot_player := false
			for {

				if bot_player {
					break
				}

				bot_action, _ := computer.PlayGame()

				if bot_action {
					card, _ := deck.GetCard()
					computer.AddCardToHand(card)
				} else {
					bot_player = true
				}
			}

			fmt.Print(bets)
			// fmt.Print("\x1B[2J\x1B[1;1H")
			break
		} else {
			break
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

func (b Blackjack) bet(player *player.Player, computer *player.Player) (Bets, error) {

	reader := bufio.NewReader(os.Stdin)
	bet_player := 0.0
	bet_computer := 0.0

	for {
		if bet_player <= 0.0 {
			fmt.Print("Insert how many will you bet in this match: ")
			capture_bet_player, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return Bets{}, nil
			}

			capture_bet_player_float, _ := strconv.ParseFloat(strings.TrimSpace(capture_bet_player), 64)
			bet_player = capture_bet_player_float
		} else {
			break
		}
	}

	_, err := player.Withdraw(bet_player)
	if err != nil {
		panic("Er")
	}

	if balance, _ := computer.GetBalance(); bet_player > balance {
		name, _ := player.GetName()
		fmt.Printf("%s, I haven't this money, I'll give all win, right?", name)
		bet_computer, _ := computer.GetBalance()
		_, err := computer.Withdraw(bet_computer)
		if err != nil {
			panic("Er")
		}

	} else {
		bet_computer = bet_player
		_, err := computer.Withdraw(bet_player)
		if err != nil {
			panic("Er")
		}
	}

	return Bets{player_bet_amount: bet_player, computer_bet_amount: bet_computer}, nil
}
