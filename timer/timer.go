package timer

import (
	"io"
	"time"
)

type Timer struct {
	WorkDuration time.Duration
	RestDuration time.Duration
}

// Start begins the timer using the set work duration.
// This is the assumed "beginning" of the timer session ie. fresh timer
func (t *Timer) Start(out io.Writer) {
	workTime := t.WorkDuration
	ticker := time.NewTicker(1 * time.Second)

	complete := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-complete:
				return
			case datetime := <-ticker.C:
				out.Write([]byte(datetime.String()))
				workTime -= time.Second
				if workTime <= 0 {
					complete <- true
				}
			}
		}
	}()
	<-complete
	ticker.Stop()
}
