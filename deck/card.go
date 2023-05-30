package deck

import "fmt"

type Card struct {
	name   string
	height []int
	ticker string
}

func (c Card) New(name string, height []int, ticker string) (Card, error) {
	card := Card{
		name:   name,
		height: height,
		ticker: ticker,
	}

	return card, nil
}

func (c Card) GetName() (string, error) {
	return c.name, nil
}

func (c Card) GetHeight() ([]int, error) {
	return c.height, nil
}

func (c Card) GetTicker() (string, error) {
	return c.ticker, nil
}

func (c Card) ShowCard() {

	if c.name == "Ten" {
		fmt.Println("- - - -")
		fmt.Println("|     |")
		fmt.Println("| {}  |", c.ticker)
		fmt.Println("|     |")
		fmt.Println("- - - -")
	} else {
		fmt.Println("- - - -")
		fmt.Println("|     |")
		fmt.Println("|  {}  |", c.ticker)
		fmt.Println("|     |")
		fmt.Println("- - - -")
	}
}
