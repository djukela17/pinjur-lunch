package models

import (
	"fmt"
	"strconv"
)

type Order struct {
	Username     string
	ChosenDish   Dish
	OptionalNote string
}

type Orders struct {
	Orders []Order
}

func (o *Orders) GetOrders() []Order {
	return o.Orders
}

func (o *Orders) AddDish(dish Dish, name, optionalNote string) {
	choice := Order{Username: name, ChosenDish: dish, OptionalNote: optionalNote}
	o.Orders = append(o.Orders, choice)
	fmt.Println(o.Orders)
}

func (o *Orders) CalcTotalPrice() int {
	total := 0
	for _, c := range o.Orders {
		total += c.ChosenDish.Price
	}
	return total
}

func (o *Orders) CreateCompressedList() []string {
	var ret []string
	stacked := make(map[string]int)

	for _, order := range o.Orders {
		fullOrder := order.ChosenDish.Name
		if order.OptionalNote != "" {
			fullOrder += "(" + order.OptionalNote + ")"
		}
		stacked[fullOrder] += 1
	}

	for k, v := range stacked {
		ret = append(ret, strconv.Itoa(v)+"x "+k)
	}
	fmt.Println(stacked)

	return ret
}
