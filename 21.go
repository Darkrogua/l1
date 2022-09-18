package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

// Имеется корзина товаров, которая содержит цены в USD (учитывает куча нюансов, как и что нам неизвестно)
// нам необходимо получать цену в RUB

type CartRub interface {
	getCartTotal() float64
}

type Cart struct {
	total float64
}

func (c *Cart) getTotal() float64 {
	return c.total
}

type AdapterCart struct {
	c *Cart
}

func (adapterCart AdapterCart) getCartTotal() float64 {
	return (adapterCart.c.total * 60.0)
}

var cart = Cart{total: 543}

func main() {
	var adapterCart CartRub = AdapterCart{c: &cart}
	fmt.Printf("Корзина содержит товары на сумму: %.0f рублей", adapterCart.getCartTotal())
}
