package main

var (
	vip_discount_grade float32 = 0.05
	Point_ratio        float32 = 1
)

type consumer struct {
	vip_level       int
	available_point float32
}

func (c consumer) normal_pay(price float32) float32 {
	return price
}

func (c consumer) vip_pay(price float32) float32 {
	price = price * (1 - (float32(c.vip_level) * vip_discount_grade))
	return price
}

func (c consumer) point_pay(price float32) float32 {
	price = price - (c.available_point / Point_ratio)
	return price
}

func (c consumer) vip_point_pay(price float32) float32 {
	if c.available_point >= 100 && c.vip_level > 0 {
		price = (price - (c.available_point / Point_ratio)) * 0.9
	} else {
		price = c.point_pay(price)
	}
	return price
}
