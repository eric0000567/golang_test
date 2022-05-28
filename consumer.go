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
func (c *consumer) add_products(item product, amount int) {
	(*c).shopping_cart[item] += amount
}

/*
add func about consumer
*/
