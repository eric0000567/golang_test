package main

import "fmt"

var payment_method = map[string]string{
	"A": "normal",
	"B": "vip",
	"C": "point",
	"D": "vip_point"}

// create product
var (
	apple   = product{"apple", 100, 1000}
	banana  = product{"banana", 80, 1000}
	jewelry = product{"jewelry", 1000, 1}
)

//create consumer
var (
	consumer1 = consumer{"Eric", 1, 1000, 50, map[product]int{}}
	consumer2 = consumer{"Alan", 2, 3000, 300, map[product]int{}}
	consumer3 = consumer{"Mark", 0, 1000, 100, map[product]int{}}
)

//create cashier
var (
	total_expenses float32 = 0
	spend_point    float32 = 0
	cashier1               = cashier{}
)

func main() {
	consumer1.add_products(&apple, 5)
	consumer1.add_products(&banana, 5)
	consumer2.add_products(&jewelry, 1)
	consumer3.add_products(&jewelry, 1)

	payment_process(&consumer1, 0, cashier1, payment_method["A"])
	payment_process(&consumer2, 150, cashier1, payment_method["D"])
	payment_process(&consumer3, 50, cashier1, payment_method["B"])
}

func payment_process(buyer *consumer, point float32, ch cashier, payment_method string) {
	total_expenses = ch.count_expenses(buyer.shopping_cart)
	if total_expenses < 0 {
		fmt.Println("Sorry, some product didn't have enough stock.")
	}
	spend_point = 0
	switch payment_method {
	case "normal":
		total_expenses = ch.normal_pay(buyer, &total_expenses)
	case "vip":
		total_expenses = ch.vip_pay(buyer, &total_expenses)
	case "point":
		total_expenses = ch.point_pay(buyer, point, &total_expenses)
		spend_point = point
	case "vip_point":
		total_expenses = ch.vip_point_pay(buyer, point, &total_expenses)
		spend_point = point
	}

	if total_expenses >= 0 {
		fmt.Println("Payment successful！ You spent $", total_expenses, " and point:", spend_point)
		ch.deduct_inventory(&buyer.shopping_cart)
	} else {
		fmt.Println("Payment fail！ please check your balance.")
	}
	(*buyer).print_balance()
}
