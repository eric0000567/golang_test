package main

import "fmt"

type consumer struct {
	name            string
	vip_level       int
	available_coin  float32
	available_point float32
	shopping_cart   map[product]int
}

func (c *consumer) print_balance() {
	fmt.Println(c.name, " vip level :", c.vip_level)
	fmt.Println(c.name, " available coin $", c.available_coin)
	fmt.Println(c.name, " available point", c.available_point)
}
func (c *consumer) add_products(item *product, amount int) {
	if (*item).available_amount < amount {
		fmt.Println("Sorry, this product don't have enough stock.")
	} else {
		(*c).shopping_cart[*item] += amount
	}
}

func (c *consumer) stored_coin(amount float32) {
	(*c).available_coin += amount
}

func (c *consumer) stored_point(amount float32) {
	(*c).available_point += amount
}

/*
add func about consumer
*/
