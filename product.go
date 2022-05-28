package main

type product struct {
	name             string
	price            float32
	available_amount int
}

func (p product) discount(ratio float32) {
	p.price = p.price * ratio
}

func (p product) supplement(amount int) {
	p.available_amount = p.available_amount + amount
}

/*
	add func about product here
*/
