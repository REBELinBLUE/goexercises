package main

import "fmt"

type Developer interface {
	Speak() string
}

type Developers []Developer

type Java struct {
	Name string
	Age  int
	Says string
}

func (p Java) Speak() string {
	return p.Says
}

type Go struct {
	Name string
	Age  int
	Says string
}

func (p Go) Speak() string {
	return p.Says
}

type Php struct {
	Name   string
	Age    int
	Shouts string
}

func (p Php) Speak() string {
	return p.Shouts
}

type C struct {
	Name     string
	Age      int
	Explains string
}

func (p C) Speak() string {
	return p.Explains
}

func main() {
	p := Developers{
		Php{Name: "connor", Age: 22, Shouts: "I like things slow and loosely typed"},
		Go{Name: "charlie", Age: 24, Says: "Throw me a Gopher"},
		Java{Name: "alex", Age: 25, Says: "I have a factory full of problems"},
		Go{Name: "bill", Age: 55, Says: "Dont talk to me about dependency management"},
		C{Name: "Jill", Age: 45, Explains: "I wish my life was simpler"},
		C{Name: "Ahamed", Age: 33, Explains: "Im tradtional and so is my language"},
		Java{Name: "Rob", Age: 55, Says: "Your only using up 99% of your resources..."},
		Php{Name: "Kimbley", Age: 23, Shouts: "Tabs and spaces lets spend loads of time debating it..."},
	}

	for _, developer := range p {
		fmt.Println(developer.Speak())
	}
}
