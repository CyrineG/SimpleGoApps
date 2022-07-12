package main

import "fmt"

func main(){
	cards := newDeckFromFile("my_cards")
	cards.print()

	cards.shuffle()
	fmt.Println("------shuffled------")
	cards.print()
}
