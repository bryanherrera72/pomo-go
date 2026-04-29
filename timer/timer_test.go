package timer

import (
	"bytes"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	t.Run("Can set up a default configuration for a timer.", func(t *testing.T) {
		got := NewConfig()
		want := config{
			WorkDuration:     25 * standardTimeMeasure,
			RestDuration:     5 * standardTimeMeasure,
			LongRestDuration: 15 * standardTimeMeasure,
			Pomos:            4,
			GoalPomos:        4,
			tickSpeed: 		  standardTickInterval,
		}

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("can customize an existing config", func(t *testing.T) {
		got := NewConfig()
		got.WorkDuration = 50 * time.Minute
		got.RestDuration = 10 * time.Minute
		got.LongRestDuration = 30 * time.Minute
		got.Pomos = 4
		got.GoalPomos = 5

		want := config{
			WorkDuration:     50 * standardTimeMeasure,
			RestDuration:     10 * standardTimeMeasure,
			LongRestDuration: 30 * standardTimeMeasure,
			Pomos:            4,
			GoalPomos:        5,
			tickSpeed: 		  standardTickInterval,
		}

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}

	})

}
func TestPomoTimer(t *testing.T) {
	t.Run("Can configure a timer with default settings.", func(t *testing.T) {
		config := NewTestConfig()
		timer := NewTimer(config)
		want := true
		// Check resulting string via byte buffer
		buff := bytes.NewBuffer(make([]byte, 10))
		timer.Start(buff)
		time.Sleep(2 * time.Second)
		got := timer.completed
		
		println(buff.String())
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}

	})

	// t.Run("Can start a fresh timer", func(t *testing.T) { //we may need a mock timer here. the regular timer passage is too long.
	// 	//here we set a timer up that doesn't follow the normal config setup.
	// 	timer := Timer{
	// 		WorkDuration:     25 * time.Millisecond,
	// 		RestDuration:     5 * time.Millisecond,
	// 		LongRestDuration: 15 * time.Millisecond,
	// 		Pomos:            4,
	// 		GoalPomos:        4,
	// 		tickSpeed:        1 * time.Millisecond,
	// 		running:          false,
	// 		completed:        false,
	// 	}
	// 	timer.Start(os.Stdout)

	// })
	// t.Run("Can configure and start a timer", func(t *testing.T) {
	// 	workInterval := 3 * time.Second // in seconds
	// 	restInterval := 3 * time.Second // in seconds
	// 	ticker := time.NewTicker(time.Millisecond)
	// 	want := 3
	// 	tickSpy := TickSpy{0}
	// 	timer := PomoTimer{workInterval, restInterval, ticker}
	// 	timer.Start(&tickSpy)
	// 	got := tickSpy.TickCount
	// 	if got != want {
	// 		t.Errorf("got %d, want: %d", got, want)
	// 	}
	// })

	// t.Run("Can reset the timer", func(t *testing.T) {
	// 	//test content.
	// })
}
