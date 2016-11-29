package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Developers working on the 5th Floor\n--------")
	fmt.Println(peopleOnFloorFive())

	fmt.Println("\nPeople managed by Charlie\n--------")
	fmt.Println(charliesTeam())

	fmt.Println("\nPeople working in Product\n--------")
	fmt.Println(peopleWorkingInProduct())
}

func peopleOnFloorFive() []string {
	fifthfloor := make([]string, 0)
	for name, details := range People {
		if details.floor == 5 && strings.ToLower(details.dept) == "development" {
			fifthfloor = append(fifthfloor, name)
		}
	}
	return fifthfloor
}

func charliesTeam() []string {
	charliesteam := make([]string, 0)
	for name, details := range People {
		if strings.ToLower(details.manager) == "charlie" {
			charliesteam = append(charliesteam, name)
		}
	}
	return charliesteam
}

func peopleWorkingInProduct() []string {
	productowners := make([]string, 0)
	for name, details := range People {
		if strings.ToLower(details.dept) == "product" {
			productowners = append(productowners, name)
		}
	}
	return productowners
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
