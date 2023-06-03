package main

import "fmt"

type Card struct {
	Name   string
	Height []int
	Ticker string
}

func (c Card) ShowCard() {

	if c.Name == "Ten" {
		fmt.Println("- - - -")
		fmt.Println("|     |")
		fmt.Printf("| %v  |\n", c.Ticker)
		fmt.Println("|     |")
		fmt.Println("- - - -")
	} else {
		fmt.Println("- - - -")
		fmt.Println("|     |")
		fmt.Printf("|  %v  |\n", c.Ticker)
		fmt.Println("|     |")
		fmt.Println("- - - -")
	}
}
