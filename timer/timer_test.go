package timer

import (
	"bytes"
	"sync"
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
		control := make(chan string)
		defer close(control)
		var wg sync.WaitGroup
		wg.Add(1)
		go timer.Time(control, buff, &wg)
		timer.Start(buff, control)
		wg.Wait()
		
		got := buff.String()
		
		if got != want {
			t.Errorf("got: '%v', want: '%v'", got, want)
		}

	})

	//NOTE: This test is finnicky since the timer may be stopped just before 2ms 
	// or just after. Will need to think of a better test method. Might only be able to 
	// reliably check the completed / running state of the timer for now. 
	t.Run("Can stop a running timer", func(t *testing.T) {
		config := NewTestConfig()
		timer := NewTimer(config)
		//testing the timer ticks with a rapid timer.
		want := "25ms24ms"
		buff := new(bytes.Buffer)
		control := make(chan string)
		defer close(control)
		var wg sync.WaitGroup
		wg.Add(1)
		go timer.Time(control, buff, &wg)
		timer.Start(buff, control)
		time.Sleep(2 * time.Millisecond)
		timer.Stop(buff, control)
		wg.Wait()
		
		got := buff.String()
		
		if got != want {
			t.Errorf("got: '%v', want: '%v'", got, want)
		}
	})

}
