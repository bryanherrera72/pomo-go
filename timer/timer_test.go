package timer

import (
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	t.Run("Can set up a default configuration for a timer.", func(t *testing.T) {
		got := NewConfig()
		want := config{
			WorkDuration:     25 * time.Minute,
			RestDuration:     5 * time.Minute,
			LongRestDuration: 15 * time.Minute,
			Pomos:            4,
			GoalPomos:        4,
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
			WorkDuration:     50 * time.Minute,
			RestDuration:     10 * time.Minute,
			LongRestDuration: 30 * time.Minute,
			Pomos:            4,
			GoalPomos:        5,
		}

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}

	})

}
func TestTimer(t *testing.T) {
	t.Run("Can configure a timer with default settings.", func(t *testing.T) {
		config := NewConfig()
		got := NewTimer(config)

		want := Timer{
			WorkDuration:     25 * time.Minute,
			RestDuration:     5 * time.Minute,
			LongRestDuration: 15 * time.Minute,
			Pomos:            4,
			GoalPomos:        4,
			tickSpeed:        1 * time.Second,
			running:          false,
			completed:        false,
		}

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}

	})

	t.Run("Can start a fresh timer", func(t *testing.T) { //we may need a mock timer here. the regular timer passage is too long.

	})
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
