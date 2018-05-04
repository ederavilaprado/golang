package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("A start...")
		time.Sleep(time.Second * 10)
		fmt.Println("A done")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("B start...")
		time.Sleep(time.Second * 8)
		fmt.Println("B done")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("C start...")
		time.Sleep(time.Second * 6)
		fmt.Println("C done")
	}()

	fmt.Println("Let's wait...")
	wg.Wait()
	fmt.Println("That's it, byee")
}
