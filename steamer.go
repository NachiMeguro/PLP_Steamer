package main

import (
	"bufio" // read input
    "fmt" // print
    "os" 
    "strconv"
    "strings"
    "sync" // for concurrency
    "time" // handle delays and durations
)

var buildTS string

// Item struct definition
type Item struct {
	Dish string
	CookTime time.durations // in seconds
}

// Map (key-value pair) of items
var items = map[string]Item{
	"Xiao Long Bao": {"Xiao Long Bao", 10 * time.second},
	"Spicy Tofu Bao": {"Spicy Tofu Bao", 7 * time.second},
	"HarGao": {"HarGao", 7 * time.second},
	"SiewMai": {"SiewMai", 9 * time.second},
	"Vegetarian Steamed Dumplin": {"Vegetarian Steamed Dumplin", 6 * time.second},
	"Char Siew Bao": {"Char Siew Bao", 6 * time.second},
	"Vegetarian Char Siew Bao": {"Vegetarian Char Siew Bao", 7 * time.second},
	"Bao": {"Bao", 5 * time.second},
	"Bao with Honey BBQ Chicken": {"Bao with Honey BBQ Chicken", 4 * time.second},
	"Bao with Pork Belly": {"Bao with Pork Belly", 4 * time.second},
	"Custard Bun": {"Custard Bun", 7 * time.second},
	"Lotus Paste Bun": {"Lotus Paste Bun", 7 * time.second},
	"Braised Pork Belly": {"Braised Pork Belly", 4 * time.second},
	"Chicken Feet Black Bean Sauce": {"Chicken Feet Black Bean Sauce", 9 * time.second}
}

var orderCounter int = 0 // to create order number

func main(){
	// See how fast the runtime from compilation
	if ts, err := 
}