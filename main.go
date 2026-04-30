package main

import (
	"os"
	"pomogo/timer"
	"time"
)

func main() {
	config := timer.NewTestConfig()
	config.WorkDuration = 500 * time.Millisecond
	timer := timer.NewTimer(config)
	
	timer.Start(os.Stdout)
}
