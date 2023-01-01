package main

import (
	"fmt"
	"time"
)

func cook(food string) {
	// Sleep pauses the current goroutine for at least the duration d.
	//  A negative or zero duration causes Sleep to return immediately.
	fmt.Println("cooking", food)
	time.Sleep(2 * time.Second)
	fmt.Println("done cooking", food)
}

func main() {
	// Now returns the current local time.
	started := time.Now()
	food := []string{"mushroom", "cake", "pizza", "pasta"}
	for _, v := range food {
		cook(v)
	}
	// Since returns the time elapsed since t. 
	fmt.Println("done in ", time.Since(started))
}
