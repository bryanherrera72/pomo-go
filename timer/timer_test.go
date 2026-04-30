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
	t.Run("Can configure and start a timer with test config.", func(t *testing.T) {
		config := NewTestConfig()
		timer := NewTimer(config)
		//testing the timer ticks with a rapid timer.
		// WARN MAY BE INCONSISTENT. May display 0s at end or not. 
		want := "25ms24ms23ms22ms21ms20ms19ms18ms17ms16ms15ms14ms13ms12ms11ms10ms9ms8ms7ms6ms5ms4ms3ms2ms1ms"
		// Check resulting string via byte buffer
		buff := new(bytes.Buffer)
		timer.Start(buff)
		got := buff.String()
		
		if got != want {
			t.Errorf("got: '%v', want: '%v'", got, want)
		}

	})

	//NOTE: need to handle states within the time. 
	t.Run("Can stop a running timer", func(t *testing.T) {
		

	})

}
