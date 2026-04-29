package timer

import (
	"io"
	"time"
)

type Timer interface {
 NewTimer(config config)
// 	Start(io.Writer)
// 	Resume(io.Writer)
// 	Stop(io.Writer)
}

type PomoTimer struct {
	WorkDuration     time.Duration
	RestDuration     time.Duration
	LongRestDuration time.Duration
	Pomos            int
	GoalPomos        int
	tickSpeed        time.Duration
	running          bool
	completed        bool
}

//NewPomoTimer: returns an instance of our base pomodorro timer. 
func NewTimer(config config) PomoTimer {
	timer := PomoTimer{}

	timer.WorkDuration = config.WorkDuration
	timer.RestDuration = config.RestDuration
	timer.LongRestDuration = config.LongRestDuration
	timer.Pomos = config.Pomos
	timer.GoalPomos = config.GoalPomos
	timer.tickSpeed = config.tickSpeed
	timer.running = false
	timer.completed = false

	return timer
}

// Start begins the timer using the set work duration.
// This is the assumed "beginning" of the timer session ie. fresh timer
func (p *PomoTimer) Start(out io.Writer) {
	// workTime := p.WorkDuration
	ticker := time.NewTicker(p.tickSpeed)

	complete := make(chan bool, 1)

	// go func() {
	// 	for {
	// 		select {
	// 		case <-complete:
	// 			return
	// 		case datetime := <-ticker.C:
	// 			out.Write([]byte(datetime.String()))
	// 			workTime -= standardTickInterval
	// 			if workTime <= 0 {
	// 				complete <- true
	// 			}
	// 		}
	// 	}
	// }()
	<-complete
	ticker.Stop()
}

//this is our worker for the timer. 
func worker(p PomoTimer, cmp chan bool, out io.Writer, workTime time.Duration){
	
	ticker := time.NewTicker(p.tickSpeed)
	for {
		select {
		case <-cmp:
			return
		case datetime := <-ticker.C:
			out.Write([]byte(datetime.String()))
			workTime -= p.tickSpeed
			if workTime <= 0 {
				cmp <- true
			}
		}
	}
	
}