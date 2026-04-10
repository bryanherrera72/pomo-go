package timer

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	// this timer type will only support setting in minutes.
	// What is a 30s work interval?
	// though seconds can be useful since we can just decrement it on a ticker.
	t.Run("Can configure and start a timer", func(t *testing.T) {
		workInterval := 3 * time.Minute // in seconds
		restInterval := 3 * time.Minute // in seconds
		buff := bytes.Buffer{}
		got := Timer{workInterval, restInterval}

		got.Start(&buff)
		fmt.Print(buff.String())
	})

	t.Run("Can reset the timer", func(t *testing.T) {
		//test content.
	})
}
