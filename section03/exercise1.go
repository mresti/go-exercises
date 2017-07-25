package section03

import (
	"fmt"
	"log"
)

// Price is the cost of something in US cents.
type Price int64

// String is the string representation of a Price
// These should be represented in US Dollars
// Example: 2595 cents => $25.95
func (p Price) String() string {
	parsed := float64(p) / 100.0
	return fmt.Sprintf("$%g", parsed)
}

// Prices is a map from an item to its price.
var Prices = map[string]Price{
	"eggs":          219,
	"bread":         199,
	"milk":          295,
	"peanut butter": 445,
	"chocolate":     150,
}

// RegisterItem adds the new item in the prices map.
// If the item is already in the prices map, a warning should be displayed to the user,
// but the value should be overwritten.
// Bonus (1pt) - Use the "log" package to print the error to the user
func RegisterItem(prices map[string]Price, item string, price Price) {
	// show log if price[item] is in prices
	if _, ok := prices[item]; ok {
		log.Print("Warning: the item is already in the prices map")
	}
	prices[item] = price
}

// Cart is a struct representing a shopping cart of items.
type Cart struct {
	Items      []string
	TotalPrice Price
}

// HasMilk returns whether the shopping cart has "milk".
func (c *Cart) HasMilk() bool {
	result := false
	for _, item := range c.Items {
		if item == "milk" {
			result = true
		}
	}
	return result
}

// HasItem returns whether the shopping cart has the provided item name.
func (c *Cart) HasItem(item string) bool {
	result := false
	for _, itemName := range c.Items {
		if itemName == item {
			result = true
		}
	}
	return result
}

// AddItem adds the provided item to the cart and update the cart balance.
// If item is not found in the prices map, then do not add it and print an error.
// Bonus (1pt) - Use the "log" package to print the error to the user
func (c *Cart) AddItem(item string) {
	// Show error
	if _, ok := Prices[item]; !ok {
		log.Print("Warning: the item is not found in the prices map")
	} else {
		c.Items = append(c.Items, item)
		c.TotalPrice += Prices[item]
	}
}

// Checkout displays the final cart balance and clears the cart completely.
func (c *Cart) Checkout() {
	//fmt.Println("Cart balance is %v", c.TotalPrice)
	log.Printf("Cart balance is %v", c.TotalPrice)
	var init Price
	c.Items = nil
	c.TotalPrice = init
}
