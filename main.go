package main

import (
	"fmt"
	"pomogo/timer"
)

func main() {

	resultString := timer.Timer(3, 3)

	fmt.Println(resultString)
	// seconds := 10
	// timer := time.NewTimer(time.Duration(seconds) * time.Second)

	// ticker := time.NewTicker(1 * time.Second)

	// fmt.Printf("Countdown: %d seconds \n", seconds)

	// go func() {
	// 	for {
	// 		select {
	// 		case <-timer.C:
	// 			fmt.Println("\nTime's up!")
	// 			return
	// 		case <-ticker.C:
	// 			seconds--
	// 			fmt.Printf("\nRemaining: %d seconds", seconds)
	// 		}
	// 	}
	// }()
}
