# golang_test

This's an easy toll system. This package is divided into consumer, project and cashier.

## Describe
Consumers have two payment methods: coins and points

Consumers can select projects and add to shopping cart

## How to use
##### 1. Create project, consumer and cashier
```go
var apple     = product{"apple", 100, 1000}
var consumer1 = consumer{"Eric", 1, 1000, 50, map[product]int{}}
var cashier1  = cashier{}
```
##### 2. Consumer bought any projects and puts it in cart
```go
consumer1.add_products(&apple, 5)
```
##### 3. Call payment_process function and give parameter
```go
payment_process(&consumer1, 0, cashier1, payment_method["A"])
```
payment_process(consumer, point, cashier, payment_method)
