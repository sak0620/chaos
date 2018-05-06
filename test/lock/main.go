package main

import (
	"fmt"
	"sync"
	"time"
)

func _main() {
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
			fmt.Print("*")
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Printf("\n%d\n", counter)
}

func __main() {
	var mu sync.Mutex
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			counter++
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println(counter)
}

func ___main() {
	var mu sync.Mutex
	wg := sync.WaitGroup{}
	counter := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			counter++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	counter := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			ch <- 1
		}()
	}

	go func() {
		for {
			i := <-ch
			counter += i
			wg.Done()
		}
	}()

	wg.Wait()
	fmt.Println(counter)
}
