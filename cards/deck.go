package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// creating a new type deck
// deck is a slice of strings

type deck []string

func newDeck() deck {
	d := deck{}
	cardSuits := []string{"spades", "hearts", "diamond", "clubs"}
	cardValues := []string{"ace","two","three","four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			d= append(d, (value +" of "+suit))
		} 
	}
	return d
}

func newCard() string {
	return "Ace of hearts"
}

func (d deck) print(){
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {

	return strings.Join([]string(d),",")
}

func (d deck) saveToFile( filename string) error{
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	
	s := strings.Split(string(bs), ",")
	return deck(s)
}