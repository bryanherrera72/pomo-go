package main

import (
	"os"
	"pomogo/timer"
	"sync"
	"time"
)

func main() {
	config := timer.NewTestConfig()
	config.WorkDuration = 500 * time.Millisecond
	timer := timer.NewTimer(config)

	//this is how the timer will behave. we'll need a wait group for the timer. 
	//and its controls
	control := make(chan string)
	defer close(control)

	var wg sync.WaitGroup
	wg.Add(1)
	go timer.Time(control, os.Stdout, &wg)
	timer.Start(os.Stdout, control)
	wg.Wait()

}