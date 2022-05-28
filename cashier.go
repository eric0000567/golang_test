package main

// init parameter
var (
	vip_discount_grade float32 = 0.05
	point_ratio        float32 = 1
)

type cashier struct {
}

func (cash cashier) normal_pay(buyer *consumer, total_expenses *float32) float32 {
	if ((*buyer).available_coin - *total_expenses) < 0 {
		return -1
	}
	(*buyer).available_coin -= *total_expenses
	return *total_expenses
}
func (cash cashier) vip_pay(buyer *consumer, total_expenses *float32) float32 {
	*total_expenses *= (1 - (float32((*buyer).vip_level) * vip_discount_grade))
	return cash.normal_pay(buyer, total_expenses)
}
func (cash cashier) point_pay(buyer *consumer, point float32, total_expenses *float32) float32 {
	if buyer.available_point-point < 0 {
		return -1
	}
	*total_expenses -= (point / point_ratio)
	if (buyer.available_coin - *total_expenses) < 0 {
		return -1
	}
	(*buyer).available_coin -= *total_expenses
	(*buyer).available_point -= point
	return *total_expenses
}
func (cash cashier) vip_point_pay(buyer *consumer, point float32, total_expenses *float32) float32 {
	if point >= 100 && buyer.vip_level > 0 {
		*total_expenses -= (point / point_ratio) * 0.9
	}
	return cash.point_pay(buyer, point, total_expenses)

}

func (cash cashier) count_expenses(items map[product]int) float32 {
	var total_expenses float32 = 0
	for item := range items {
		if item.available_amount < items[item] {
			return -1
		}
		total_expenses += item.price * float32(items[item])
	}
	return total_expenses
}

func (cash cashier) deduct_inventory(items *map[product]int) float32 {
	var total_expenses float32 = 0
	for item := range *items {
		item.available_amount -= (*items)[item]
	}
	return total_expenses
}
