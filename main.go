package main

import "fmt"

var (
	consumer1 = consumer{1}
	consumer2 = consumer{2}
	consumer3 = consumer{0}
	consumer4 = consumer{5}
)

func main() {
	fmt.Println(consumer1.normal_pay(2000))
	fmt.Println(consumer2.vip_pay(2000))
	fmt.Println(consumer3.point_pay(2000, 100))
	fmt.Println(consumer4.vip_point_pay(2000, 500))
}
