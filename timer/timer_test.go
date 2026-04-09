package timer

import (
	"fmt"
	"testing"
)

func TestTimer(t *testing.T) {
	// this timer type will only support setting in minutes.
	// What is a 30s work interval?
	t.Run("Can configure and start a timer", func(t *testing.T) {
		workInterval := 3 // in seconds
		restInterval := 3 // in seconds
		got := Timer(workInterval, restInterval)
		want := `3
2
1
Time limit reached.`
		if got != want {
			fmt.Printf("got %v, want %v", got, want)
		}
	})

	t.Run("Can reset the timer", func(t *testing.T) {
		//test content.
	})
}
