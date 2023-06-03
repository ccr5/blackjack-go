package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	balanceInFloat, err := strconv.ParseFloat(strings.TrimSpace(balance), 64)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	human := Human{
		Player: Player{
			Name:    name,
			Hand:    make([]Card, 0),
			Balance: balanceInFloat,
			Wins:    0,
			Defeats: 0,
		},
	}

	fmt.Print("Could you choose a name for a bot? ")
	computerName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	computer := Bot{
		Player: Player{
			Name:    computerName,
			Hand:    make([]Card, 0),
			Balance: balanceInFloat,
			Wins:    0,
			Defeats: 0,
		},
	}
	computer.WelcomeMessage()

	fmt.Println("So let's play? ")
	fmt.Print("\x1B[2J\x1B[1;1H")

	deck := Deck{}
	game := false

	for {
		if !game {
			human.ClearHand()
			computer.ClearHand()
			deck.CreateDeck()
			deck.ShuffleDeck()
			fmt.Print("\x1B[2J\x1B[1;1H")

			checkBalancePlayer, _ := b.checkBalance(human.Player)
			checkBalanceComputer, _ := b.checkBalance(computer.Player)

			if checkBalancePlayer || checkBalanceComputer {
				game = true
				continue
			}

			bets, err := b.bet(&human.Player, &computer.Player)
			if err != nil {
				panic("Bet error")
			}

			askPlayer := false
			for {

				if askPlayer {
					break
				}

				if len(human.Hand) == 0 {
					fmt.Println("You don't have any card in your hand")
				} else {
					human.ShowHand()
				}

				fmt.Print("Would you like one more card? ")
				moreCard, err := reader.ReadString('\n')
				if err != nil {
					panic("Error to check more card")
				}

				moreCardUpper := strings.ToUpper(strings.TrimSpace(moreCard))

				if moreCardUpper == "Y" {
					card, _ := deck.GetCard()
					human.AddCardToHand(card)
				} else {
					askPlayer = true
				}
			}

			botPlayer := false
			for {

				if botPlayer {
					break
				}

				botAction, _ := computer.PlayGame()

				if botAction {
					card, _ := deck.GetCard()
					computer.AddCardToHand(card)
				} else {
					botPlayer = true
				}
			}

			winner, verifyH1, verifyH2 := CheckWinner(human.Hand, computer.Hand)

			b.checkResult(
				&human.Player,
				&computer.Player,
				bets,
				winner,
				verifyH1,
				verifyH2,
			)

			human.ShowInfo()
			human.ShowHand()

			computer.ShowInfo()
			computer.ShowHand()

			fmt.Println("Do you wanna play again (Y/N): ")
			choice, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}

			game = b.checkPlayAgain(strings.TrimSpace(choice))
		} else {
			break
		}
	}
}

func (b Blackjack) checkBalance(player Player) (bool, error) {
	if player.Balance == 0.0 {
		fmt.Printf("%v, you haven't balance\n", player.Name)
		return true, nil
	} else {
		return false, nil
	}
}

func (b Blackjack) bet(player *Player, computer *Player) (Bets, error) {

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

	if bet_player > computer.Balance {
		fmt.Printf("%v, I haven't this money, I'll give all win, right? \n", player.Name)
		_, err := computer.Withdraw(computer.Balance)
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

	return Bets{PlayerBetAmount: bet_player, ComputerBetAmount: bet_computer}, nil
}

func (b Blackjack) checkResult(
	player *Player,
	computer *Player,
	bets Bets,
	winner string,
	total_player int,
	total_computer int,
) {
	if winner == "no win" {
		fmt.Printf("no one win, you guys have more than 21!")
		fmt.Printf("%v: %v \t%v: %v\n", player.Name, total_player, computer.Name, total_computer)
		player.Deposit(bets.PlayerBetAmount)
		computer.Deposit(bets.ComputerBetAmount)
	} else if winner == "draw" {
		fmt.Printf("it's a draw O.o")
		fmt.Printf("%v: %v \t%v: %v\n", player.Name, total_player, computer.Name, total_computer)
		player.Deposit(bets.PlayerBetAmount)
		computer.Deposit(bets.ComputerBetAmount)
		player.AddWin()
		computer.AddWin()
	} else if winner == "h1" {
		fmt.Printf("%v win !!\n", player.Name)
		fmt.Printf("%v: %v \t%v: %v\n", player.Name, total_player, computer.Name, total_computer)
		player.Deposit(bets.PlayerBetAmount + bets.ComputerBetAmount)
		player.AddWin()
		computer.AddDefeats()
	} else {
		fmt.Printf("%v win !!\n", computer.Name)
		fmt.Printf("%v: %v \t%v: %v\n", player.Name, total_player, computer.Name, total_computer)
		computer.Deposit(bets.PlayerBetAmount + bets.ComputerBetAmount)
		computer.AddWin()
		player.AddDefeats()
	}
}

func (b Blackjack) checkPlayAgain(choice string) bool {
	if strings.ToUpper(choice) == "Y" {
		return false
	} else {
		return true
	}
}
