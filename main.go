package main

import (
	"fmt"
	"os"
	"pomogo/timer"
	"time"
)

func main() {
	config := timer.NewTestConfig()
	timer := timer.NewTimer(config)
	timer.Start(os.Stdout)
	fmt.Println("value: %v", timer.WorkDuration-(25*time.Minute) == time.Second)
}
