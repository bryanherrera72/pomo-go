package timer

import (
	"testing"
	"time"
)

type TickSpy struct{
	TickCount int
}

func (t *TickSpy) Write(p []byte) (n int, err error){
	t.TickCount++
	return 
}


func TestTimer(t *testing.T) {
	
	t.Run("Can configure and start a timer", func(t *testing.T) {
		workInterval := 3 * time.Second // in seconds
		restInterval := 3 * time.Second // in seconds
		want := 3
		tickSpy := TickSpy{0}
		timer := Timer{workInterval, restInterval}
		timer.Start(&tickSpy)
		got := tickSpy.TickCount
		if(got != want){
			t.Errorf("got %d, want: %d", got, want)
		}
	})

	t.Run("Can reset the timer", func(t *testing.T) {
		//test content.
	})
}
