package main

import (
	"fmt"
	"sync"
)

type Container struct {
	sync.Mutex
	counter map[string]int
}

func (c *Container) inc(name string) {
	c.Lock()
	defer c.Unlock()
	c.counter[name]++
}

func (c *Container) doIncrement(wg *sync.WaitGroup, name string, n int) {
	for i := 0; i < n; i++ {
		c.inc(name)
	}
	wg.Done()
}

func main() {
	c := Container{
		counter: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup

	wg.Add(3)
	go c.doIncrement(&wg, "a", 10000)
	go c.doIncrement(&wg, "a", 10000)
	go c.doIncrement(&wg, "b", 10000)
	wg.Wait()
	fmt.Println(c.counter)
}
