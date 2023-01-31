package main

import "fmt"

type deck []string

func (d deck) printCards() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
