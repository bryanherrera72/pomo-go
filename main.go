package main

import (
	"bytes"
	"fmt"
	"pomogo/timer"
	"time"
)

func main() {
	workInterval := 3 * time.Second // in seconds
	restInterval := 3 * time.Second // in seconds

	buff := bytes.Buffer{}
	myTimer := timer.Timer{
		WorkDuration: workInterval,
		RestDuration: restInterval}
	myTimer.Start(&buff)
	fmt.Println(buff.String())
	
}
