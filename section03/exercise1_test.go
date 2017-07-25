package section03

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"
)

func TestPrice_String(t *testing.T) {
	var input Price = 2595
	want := "$25.95"
	result := input.String()
	if result != want {
		t.Error(`2595.String() != $25.95`)
	}
}

func TestPrice_StringFull(t *testing.T) {
	var tests = []struct {
		input Price
		want  string
	}{
		{2595, "$25.95"},
		{234545, "$2345.45"},
		{3, "$0.03"},
	}
	for _, test := range tests {
		if got := test.input.String(); got != test.want {
			t.Errorf("%v.String() = %v", test.input, got)
		}
	}
}

func TestRegisterItem_NewItem(t *testing.T) {
	var pricesInput = map[string]Price{
		"eggs": 219,
	}

	var pricesWant = map[string]Price{
		"eggs":  219,
		"water": 259,
	}
	var input Price = 259
	item := "water"

	RegisterItem(pricesInput, item, input)
	if !reflect.DeepEqual(pricesInput, pricesWant) {
		t.Error(`RegisterItem(map[eggs:$2.19],water,$2.59) = map[water:$2.59 eggs:$2.19]`)
	}
}

func TestRegisterItem_ExistItem(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	var pricesInput = map[string]Price{
		"eggs": 219,
	}

	var pricesWant = map[string]Price{
		"eggs": 259,
	}
	var input Price = 259
	item := "eggs"
	msgExpected := "Warning: the item is already in the prices map"

	RegisterItem(pricesInput, item, input)
	msgReceived := buf.String()

	// Check value edited
	if !reflect.DeepEqual(pricesInput, pricesWant) {
		t.Error(`RegisterItem(map[eggs:$2.19],eggs,$2.59) = map[eggs:$2.59]`)
	}
	// Check log message
	if !strings.HasSuffix(msgReceived[:len(msgReceived)-1], msgExpected) {
		t.Fatalf("Unexpected log message. Got %s but should contain %s", msgReceived[:len(msgReceived)-1], msgExpected)
	}
}

func TestCart_HasMilk(t *testing.T) {
	input := Cart{[]string{"milk", "apple", "orange"}, Price(21)}
	if !input.HasMilk() {
		t.Error(`Cart{[]string{"milk","apple","oreange"}, Price(11)}.HasMilk() != true`)
	}
}

func TestCart_HasMilkFull(t *testing.T) {
	var tests = []struct {
		input Cart
		want  bool
	}{
		{Cart{[]string{"apple", "orange", "milk"}, Price(11)}, true},
		{Cart{[]string{"water", "apple"}, Price(22)}, false},
	}
	for _, test := range tests {
		if got := test.input.HasMilk(); got != test.want {
			t.Errorf("%v.HasMilk() = %v", test.input, got)
		}
	}
}

func TestCart_HasItem(t *testing.T) {
	input := Cart{[]string{"milk", "apple", "orange"}, Price(224)}
	itemName := "apple"
	if !input.HasItem(itemName) {
		t.Error(`Cart{[]string{"milk","apple","oreange"}, Price(224)}.HasItem("apple") != true`)
	}
}

func TestCart_HasItemFull(t *testing.T) {
	var tests = []struct {
		input    Cart
		itemName string
		want     bool
	}{
		{Cart{[]string{"apple", "orange", "milk"}, Price(224)}, "water", false},
		{Cart{[]string{"water", "apple"}, Price(112)}, "apple", true},
	}
	for _, test := range tests {
		if got := test.input.HasItem(test.itemName); got != test.want {
			t.Errorf("%v.HasItem(%v) = %v", test.input, test.itemName, got)
		}
	}
}

func TestCart_AddItem(t *testing.T) {
	var input Cart
	want := Cart{[]string{"milk"}, Price(295)}
	input.AddItem("milk")
	if !reflect.DeepEqual(input, want) {
		t.Errorf("%v != %v", input, want)
	}
}

func TestCart_AddItem_NotExistItemPrice(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	var input Cart
	var want Cart
	item := "apple"
	msgExpected := "Warning: the item is not found in the prices map"

	input.AddItem(item)
	msgReceived := buf.String()

	// Check add item that it is not exist item in Prices
	if !reflect.DeepEqual(input, want) {
		t.Error(`{[] $0.00}.AddItem("apple") = {[] $0.00}`)
	}

	// Check log message
	if !strings.HasSuffix(msgReceived[:len(msgReceived)-1], msgExpected) {
		t.Fatalf("Unexpected log message. Got %s but should contain %s", msgReceived[:len(msgReceived)-1], msgExpected)
	}
}

func TestCart_Checkout(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	var input Cart
	finalCart := Cart{[]string{"milk", "eggs"}, Price(514)}
	var want Cart
	priceExpected := Price(514)
	msgExpected := fmt.Sprintf("Cart balance is %v", priceExpected)

	// Add items
	input.AddItem("milk")
	input.AddItem("eggs")

	// Check that input have items
	if !reflect.DeepEqual(input, finalCart) {
		t.Error(`{[] $0.00}.AddItem("milk") = {[milk] $2.95}`)
		t.Error(`{[milk] $2.95}.AddItem("eggs") = {[milk eggs] $5.14}`)
	}

	// Display cart
	input.Checkout()
	msgReceived := buf.String()

	// Check the cart is empty
	if !reflect.DeepEqual(input, want) {
		t.Error(`{[] $0.00}.Checkout() = {[] $0.00}`)
	}

	// Check log message
	if !strings.HasSuffix(msgReceived[:len(msgReceived)-1], msgExpected) {
		t.Fatalf("Unexpected log message. Got %s but should contain %s", msgReceived[:len(msgReceived)-1], msgExpected)
	}
}
