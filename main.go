package main

import "fmt"

var (
	consumer1 = consumer{1, 100}
	consumer2 = consumer{2, 300}
	consumer3 = consumer{0, 100}
	consumer4 = consumer{5, 900}
)

func main() {
	fmt.Println("you need spend $", consumer1.normal_pay(2000))
	fmt.Println("you need spend $", consumer2.vip_pay(2000))
	fmt.Println("you need spend $", consumer3.point_pay(2000))
	fmt.Println("you need spend $", consumer4.vip_point_pay(2000))
}
