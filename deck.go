package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	var cards deck

	cardSuits := []string{"Heart", "Spades", "Diamond", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, s := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" Of "+s)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		// Option #1 - log the error and call newDeck method and return a new Deck
		// Option #2 - log the error and exit the program
		fmt.Println("Error ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) suffle() deck {
	// Sourse of new Random
	sourse := rand.NewSource(time.Now().UnixNano())
	r := rand.New(sourse)

	for i := range d {
		rndPos := r.Intn(len(d) - 1)
		// Swap the values
		d[rndPos], d[i] = d[i], d[rndPos]
	}
	return d
}
