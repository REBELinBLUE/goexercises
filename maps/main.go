package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Developers working on the 5th Floor\n--------")
	for name, details := range People {
		if details.floor == 5 && strings.ToLower(details.dept) == "development" {
			fmt.Println(name)
		}
	}

	fmt.Println("\nPeople managed by Charlie\n--------")
	for name, details := range People {
		if strings.ToLower(details.manager) == "charlie" {
			fmt.Println(name)
		}
	}

	fmt.Println("\nPeople working in Product\n--------")
	for name, details := range People {
		if strings.ToLower(details.dept) == "product" {
			fmt.Println(name)
		}
	}
}

type details struct {
	age     int
	dept    string
	floor   int
	manager string
}

var People = map[string]details{
	"charlie": details{
		age:     22,
		dept:    "development",
		manager: "Fraser",
		floor:   5,
	},
	"john": details{
		age:     34,
		dept:    "Product",
		manager: "Chris",
		floor:   5,
	},
	"bob": details{
		age:     24,
		dept:    "Development",
		manager: "Charlie",
		floor:   5,
	},
	"chris": details{
		age:     34,
		dept:    "Development",
		manager: "Fraser",
		floor:   3,
	},
	"nadim": details{
		age:     44,
		dept:    "Development",
		manager: "Fraser",
		floor:   3,
	},
	"josh": details{
		age:     25,
		dept:    "Development",
		manager: "Jai",
		floor:   5,
	},
	"micheal": details{
		age:     28,
		dept:    "Development",
		manager: "Charlie",
		floor:   5,
	},
	"alex": details{
		age:     32,
		dept:    "Development",
		manager: "Charlie",
		floor:   5,
	},
	"andrew": details{
		age:     44,
		dept:    "Development",
		manager: "Charlie",
		floor:   3,
	},
}
