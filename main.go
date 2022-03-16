package main

import (
	"fmt"
	client_one "github.com/nelsongp/singleton/client-one"
	client_two "github.com/nelsongp/singleton/client-two"
	"github.com/nelsongp/singleton/singleton"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			client_one.IncrementAge()
		}()
		go func() {
			defer wg.Done()
			client_two.IncrementAge()
		}()
	}

	wg.Wait()
	p := singleton.GetInstance()
	age := p.GetAge()
	fmt.Println(age)
}
