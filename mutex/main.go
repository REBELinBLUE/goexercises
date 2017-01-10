package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var Animals = map[string]int{
	"Dog":      0,
	"Cat":      0,
	"Rabbit":   0,
	"Goldfish": 0,
}

type Cache struct {
	sync.Mutex
	Animals map[string]int
}

func New() *Cache {
	return &Cache{
		Animals: make(map[string]int),
	}
}

func (c *Cache) Get(key string) int {
	c.Lock()
	defer c.Unlock()
	return c.Animals[key]
}

func (c *Cache) Set(key string, value int) {
	c.Lock()
	defer c.Unlock()
	c.Animals[key] = value
}

func main() {
	ca := New()
	for animal, count := range Animals {
		ca.Set(animal, count)

		for c := 0; c < 10; c++ {
			go func(cache *Cache, animal string, count int) {
				s := rand.NewSource(time.Now().UnixNano())
				r := rand.New(s)
				num := r.Intn(1000)

				cache.Set(animal, num)

				//cache.Animals[animal] = num

				fmt.Printf("%v %v set to %d\n", count, animal, num)
			}(ca, animal, c)
		}
	}

	time.Sleep(time.Second * 2)
	fmt.Printf("%v", ca)
}
