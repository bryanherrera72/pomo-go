package timer

import (
	"fmt"
	"io"
	"time"
)

type Timer struct {
	WorkDuration time.Duration
	RestDuration time.Duration
}

// Start begins the timer using the set work duration.
// This is the assumed "beginning" of the timer session ie. fresh timer
func (t *Timer) Start(writer io.Writer) {
	work := t.WorkDuration
	ticker := time.NewTicker(1 * time.Second)

	complete := make(chan bool, 1)

	fmt.Println("Let's start this shit")
	go func() {
		for {
			select {
			case <-complete:
				return
			case datetime := <-ticker.C:
				writer.Write([]byte(datetime.String()))
				work -= time.Second
				if work <= 0*time.Second {
					complete <- true
				}
			}
		}
	}()
	<-complete
	ticker.Stop()
	fmt.Println("timer done")
	fmt.Println("finished")

}
